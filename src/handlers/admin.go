package handlers

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/markdown"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
	"github.com/a-h/templ"
	"golang.org/x/oauth2"
)

type AdminHandler struct {
	discordOAuthConfig *oauth2.Config
	AdminUserID        string
	isProduction       bool
	cookieSecret       []byte
}

func NewAdminHandler(isProduction bool) *AdminHandler {
	secret := os.Getenv("COOKIE_SECRET")

	return &AdminHandler{
		discordOAuthConfig: &oauth2.Config{
			ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
			ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("REDIRECT_URL"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://discord.com/api/oauth2/authorize",
				TokenURL: "https://discord.com/api/oauth2/token",
			},
			Scopes: []string{"identify"},
		},
		AdminUserID:  "758322333437394944",
		isProduction: isProduction,
		cookieSecret: []byte(secret),
	}
}

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *AdminHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.discordOAuthConfig.AuthCodeURL("state-pseudo-random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *AdminHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	ctx := context.Background()

	token, err := h.discordOAuthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := h.discordOAuthConfig.Client(ctx, token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		http.Error(w, "Failed to fetch user profile", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user DiscordUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse user profile", http.StatusInternalServerError)
		return
	}

	if user.ID != h.AdminUserID {
		http.Error(w, "Unauthorized: You are not allowed to be here, go away.", http.StatusForbidden)
		return
	}

	var useSecure bool
	if h.isProduction {
		useSecure = true
	} else {
		useSecure = false
	}

	signedVal := h.signValue(user.ID)

	http.SetCookie(w, &http.Cookie{
		Name:     "admin_session",
		Value:    signedVal,
		Path:     "/",
		HttpOnly: true,
		Secure:   useSecure,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}

func (h *AdminHandler) HandleAdmin(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("admin_session")
	if err != nil {
		component := views.Admin(false)
		component.Render(r.Context(), w)
		return
	}

	if !h.hasAccess(sessionCookie.Value) {
		component := views.Admin(false)
		component.Render(r.Context(), w)
		return
	}

	component := views.Admin(true)
	component.Render(r.Context(), w)
}

func (h *AdminHandler) HandleAdminLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "admin_session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // deletes cookie
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}

// ---

func (h *AdminHandler) signValue(value string) string {
	mac := hmac.New(sha256.New, h.cookieSecret)
	mac.Write([]byte(value))
	signature := hex.EncodeToString(mac.Sum(nil))
	return value + "." + signature
}

func (h *AdminHandler) verifyValue(signedValue string) (string, bool) {
	parts := strings.Split(signedValue, ".")
	if len(parts) != 2 {
		return "", false
	}
	value, signatureHex := parts[0], parts[1]

	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return "", false
	}

	mac := hmac.New(sha256.New, h.cookieSecret)
	mac.Write([]byte(value))
	expectedMAC := mac.Sum(nil)

	if hmac.Equal(signature, expectedMAC) {
		return value, true
	}
	return "", false
}

func (h *AdminHandler) hasAccess(signedValue string) bool {
	userID, valid := h.verifyValue(signedValue)
	if !valid || userID != h.AdminUserID {
		return false
	}
	return true
}

// --- DIFFERENT ADMIN PAGES

func (h *AdminHandler) HandleSpecificAdminPages(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("admin_session")
	if err != nil {
		component := views.Admin(false)
		component.Render(r.Context(), w)
		return
	}

	if !h.hasAccess(sessionCookie.Value) {
		component := views.Admin(false)
		component.Render(r.Context(), w)
		return
	}

	var component templ.Component
	switch r.PathValue("specificPage") {
	case "upload-wiki":
		component = views.UploadWikiView()
	default:
		component = views.NotFound()
	}

	component.Render(r.Context(), w)
}

func (h *AdminHandler) HandleAPI(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("admin_session")
	if err != nil {
		component := views.Admin(false)
		component.Render(r.Context(), w)
		return
	}

	if !h.hasAccess(sessionCookie.Value) {
		component := views.Admin(false)
		component.Render(r.Context(), w)
		return
	}

	switch r.PathValue("specificCommand") {
	case "upload-wiki":
		h.handleUploadedWikiArticle(w, r)
	case "preview-wiki":
		h.handleWikiPreview(w, r)
	}
}

// ---

func (h *AdminHandler) handleUploadedWikiArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	prettyName := r.FormValue("pretty-name")
	categoryPath := r.FormValue("category-path")

	markdownSource := r.FormValue("md-source")
	HTML := markdown.GenerateHTMLFromString(markdownSource)

	sortOrderString := r.FormValue("sort-order")
	sortOrder, err := strconv.Atoi(sortOrderString)
	if err != nil {
		sortOrder = 999
	}

	lastEditedAt := time.Now().Unix()

	wikiArticle := myTypes.Article{
		Name:           name,
		PrettyName:     prettyName,
		HTML:           HTML,
		MarkdownSource: markdownSource,
		CategoryPath:   categoryPath,
		SortOrder:      sortOrder,
		Type:           myTypes.WikiArticle,
		Info: myTypes.ArticleInfo{
			Date: lastEditedAt,
		},
	}

	err = database.SaveWikiArticle(wikiArticle)
	if err != nil {
		fmt.Fprintf(w, "Database error: %v", err)
		return
	}

	fmt.Fprintf(w, "Successfully saved wiki article '%v'", wikiArticle.Name)
}

func (h *AdminHandler) handleWikiPreview(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	prettyName := r.FormValue("pretty-name")
	categoryPath := r.FormValue("category-path")

	markdownSource := r.FormValue("md-source")
	HTML := markdown.GenerateHTMLFromString(markdownSource)

	sortOrderString := r.FormValue("sort-order")
	sortOrder, err := strconv.Atoi(sortOrderString)
	if err != nil {
		sortOrder = 999
	}

	lastEditedAt := time.Now().Unix()

	wikiArticle := myTypes.Article{
		Name:           name,
		PrettyName:     prettyName,
		HTML:           HTML,
		MarkdownSource: markdownSource,
		CategoryPath:   categoryPath,
		SortOrder:      sortOrder,
		Type:           myTypes.WikiArticle,
		Info: myTypes.ArticleInfo{
			Date: lastEditedAt,
		},
	}

	component := views.GenericArticle(wikiArticle)
	component.Render(r.Context(), w)
}
