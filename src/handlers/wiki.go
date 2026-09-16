package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/markdown"
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

func (h *WikiHandler) findWhichArticleWasRequested(r *http.Request) (string, error) {
	if r.URL.String() == "/wiki/" {
		return "introduction", nil
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
			return splitURL[2], nil
		}
	}
}

func (h *WikiHandler) getRequestedArticle(requestedArticle string) (myTypes.Article, error) {
	html := markdown.GenerateHTMLFromString(loadTestFile())
	article := myTypes.Article{
		Name:       requestedArticle,
		PrettyName: "Test Article",
		HTML:       html,
		Info: myTypes.ArticleInfo{
			Date:   "09/03/2026",
			Author: "VOVOplay",
		},
	}

	return article, nil

}

func (h *WikiHandler) HandleWiki(w http.ResponseWriter, r *http.Request) {
	requestedArticleName, err := h.findWhichArticleWasRequested(r)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	requestedArticle, err := h.getRequestedArticle(requestedArticleName)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	component := views.WikiLayout(requestedArticle)
	component.Render(r.Context(), w)
}
