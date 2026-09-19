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

	allArticles := h.getArticles()

	// for caching
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.WikiLayout(requestedArticle, allArticles)
	component.Render(r.Context(), w)
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
	if requestedArticle != "first-article" {
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

	article := myTypes.Article{
		Name:       requestedArticle,
		PrettyName: "Test Article",
		HTML:       "<p>wow</p>",
		Info: myTypes.ArticleInfo{
			Date:   "09/03/2026",
			Author: "VOVOplay",
		},
	}
	return article, nil
}

// temp
func (h *WikiHandler) getArticles() []myTypes.Article {
	return []myTypes.Article{
		{
			Name:         "first-article",
			PrettyName:   "First Article",
			HTML:         "FIRST ARTICLE",
			CategoryPath: "/",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "second-article",
			PrettyName:   "Second Article",
			HTML:         "SECOND ARTICLE",
			CategoryPath: "/Perks/",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "third-article",
			PrettyName:   "Third Article",
			HTML:         "THIRD ARTICLE",
			CategoryPath: "/Perks/",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "fourth-article",
			PrettyName:   "Fourth Article",
			HTML:         "FOURTH ARTICLE",
			CategoryPath: "/Perks/Talent Roles",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
	}
}
