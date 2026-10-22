package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/markdown"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

var ErrBlogNotFound error = errors.New("Blog post with this name not found")

type BlogHandler struct {
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

func (h *BlogHandler) HandleSpecificBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=60")

	requestedBlogName := r.PathValue("specificBlogPost")

	if !h.helperIsValidURL(r.URL) {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	requestedBlog, err := database.GetBlogByName(requestedBlogName, markdown.GenerateHTMLFromString)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	component := views.SpecificBlogView(requestedBlog)
	component.Render(r.Context(), w)
}

func (h *BlogHandler) HandleBlogLandingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=60")

	allBlogs, err := database.GetAllBlogs()
	databaseOkay := true
	if err != nil {
		databaseOkay = false
	}

	component := views.BlogLandingPage(allBlogs, databaseOkay)
	component.Render(r.Context(), w)
}

func (h *BlogHandler) helperIsValidURL(URL *url.URL) bool {
	pathSections := strings.Split(URL.Path, "/")
	if len(pathSections) > 3 && pathSections[3] != "" {
		return false
	}
	return true
}
