package handlers

import (
	"cmp"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/markdown"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

var ErrBlogNotFound error = errors.New("Blog post with this name not found")
var ErrPageDoesntExist error = errors.New("This page does not exist")
var pageSize int = 10

type blogCache struct {
	mutex sync.Mutex
	ok    bool
	pages [][]myTypes.Article // [[article, article], [article article]]
}

func (c *blogCache) Update() {
	c.pages = [][]myTypes.Article{}

	blogs, err := database.GetAllBlogs()
	if err != nil {
		c.mutex.Lock()
		c.ok = false
		c.mutex.Unlock()

		fmt.Printf("Error in blog page cache: %v\n", err)
		return
	}

	slices.SortFunc(blogs, func(a myTypes.Article, b myTypes.Article) int {
		return cmp.Compare(b.Info.Date, a.Info.Date)
	})

	var newPages [][]myTypes.Article
	for i := 0; i < len(blogs); i += pageSize {
		end := i + pageSize
		if end > len(blogs) {
			end = len(blogs)
		}

		newPages = append(newPages, blogs[i:end])
	}

	c.mutex.Lock()
	c.pages = newPages
	c.ok = true
	c.mutex.Unlock()
}

func (c *blogCache) GetPage(pageNumber int) ([]myTypes.Article, error) {
	if pageNumber >= len(c.pages) {
		return []myTypes.Article{}, ErrPageDoesntExist
	}

	return c.pages[pageNumber], nil
}

type BlogHandler struct {
	blogCache *blogCache
}

func NewBlogHandler() *BlogHandler {
	blogCache := &blogCache{}
	blogCache.Update()

	return &BlogHandler{
		blogCache: blogCache,
	}
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

func (h *BlogHandler) LoadMore(w http.ResponseWriter, r *http.Request) {
	pageNumberString := r.URL.Query().Get("pageNumber")
	pageNumber, err := strconv.Atoi(pageNumberString)
	if err != nil {
		http.Error(w, "Page number must be an integer", http.StatusNotAcceptable)
		return
	}

	blogs, err := h.blogCache.GetPage(pageNumber)
	fmt.Println(blogs, err) // TODO - Render templ component
}
