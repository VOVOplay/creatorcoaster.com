package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

var ErrBlogNotFound error = errors.New("Blog post with this name not found")

type BlogHandler struct {
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

func (h *BlogHandler) HandleBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=60")

	if r.URL.String() == "/blog/" { // blog landing page
		allBlogs, err := database.GetAllBlogs()
		databaseOkay := true
		if err != nil {
			databaseOkay = false
		}

		component := views.BlogLandingPage(allBlogs, databaseOkay)
		component.Render(r.Context(), w)

		return
	}

	requestedBlogName, err := h.findWhichBlogWasRequested(r)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	requestedBlog, err := database.GetBlogByName(requestedBlogName)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	component := views.SpecificBlogView(requestedBlog)
	component.Render(r.Context(), w)
}

func (h *BlogHandler) findWhichBlogWasRequested(r *http.Request) (string, error) {
	URL := r.URL.String()
	splitURL := strings.Split(URL, "/")

	var cleanedSplitURL []string
	for _, URLSection := range splitURL {
		if URLSection != "" {
			cleanedSplitURL = append(cleanedSplitURL, URLSection)
		}
	}

	if len(cleanedSplitURL) != 2 { // two parts: /blog/blog-post
		return "", InvalidArticleName
	} else {
		return cleanedSplitURL[1], nil
	}
}
