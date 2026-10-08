package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

var InvalidArticleName error = errors.New("Invalid article name")
var ArticleNotFound error = errors.New("Article Not Found")

type WikiHandler struct {
}

func NewWikiHandler() *WikiHandler {
	return &WikiHandler{}
}

func (h *WikiHandler) HandleWiki(w http.ResponseWriter, r *http.Request) {
	requestedArticleName, err := h.findWhichArticleWasRequested(r)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	requestedArticle, err := database.GetWikiArticleByName(requestedArticleName)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	allArticles, err := database.GetAllWikiArticles()
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	// for caching
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.WikiLayout(requestedArticle, allArticles)
	component.Render(r.Context(), w)
}

func (h *WikiHandler) findWhichArticleWasRequested(r *http.Request) (string, error) {
	if r.URL.String() == "/wiki/" {
		return myTypes.FirstArticleName, nil
	} else {
		URL := r.URL.String()
		splitURL := strings.Split(URL, "/")

		var cleanedSplitURL []string
		for _, URLSection := range splitURL {
			if URLSection != "" {
				cleanedSplitURL = append(cleanedSplitURL, URLSection)
			}
		}

		if len(cleanedSplitURL) != 2 { // two parts: /wiki/article-name
			return "", InvalidArticleName
		} else {
			return cleanedSplitURL[1], nil
		}
	}
}
