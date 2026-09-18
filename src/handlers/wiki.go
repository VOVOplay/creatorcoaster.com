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
			CategoryPath: "/perks/",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "third-article",
			PrettyName:   "Third Article",
			HTML:         "THIRD ARTICLE",
			CategoryPath: "/perks/",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "fourth-article",
			PrettyName:   "Fourth Article",
			HTML:         "FOURTH ARTICLE",
			CategoryPath: "/perks/talent",
			Info: myTypes.ArticleInfo{
				Date:   "09/08/2027",
				Author: "VOVOplay",
			},
		},
	}
}

// Returns a slice of category names. The first being the top-most category, and the last being the lowest category.
// Example: /folder-one/folder-two/ -> ["folder-one", "folder-two"]
func formatPath(rawPath string) []string {
	splitBySlashPath := strings.Split(rawPath, "/")

	// cleans up the output of strings.Split so it only includes the path names themselves
	var formattedPath []string
	for _, pathName := range splitBySlashPath {
		if pathName == "" {
			continue
		} else {
			formattedPath = append(formattedPath, pathName)
		}
	}

	return formattedPath
}

func buildCategoryForArticle(currentLevel *[]myTypes.CategoryItem, path []string, article *myTypes.Article) {
	if len(path) == 0 {
		*currentLevel = append(*currentLevel, article)
		return
	}

	topLevelCategoryName := path[0] // top level meaning the highest for that article
	var topLevelCategory *myTypes.Category

	categoryExists := false
	for _, item := range *currentLevel {
		switch item := item.(type) {
		case *myTypes.Category:
			if item.Name == topLevelCategoryName {
				topLevelCategory = item
				categoryExists = true
				break
			}
		}
	}

	if !categoryExists {
		topLevelCategory = &myTypes.Category{
			Name:     topLevelCategoryName,
			Elements: []myTypes.CategoryItem{},
		}
		*currentLevel = append(*currentLevel, topLevelCategory) // adds it to the current level
	}

	buildCategoryForArticle(&topLevelCategory.Elements, path[1:], article)
}

func BuildCategories(allArticles []myTypes.Article) []myTypes.CategoryItem {
	var categoryTree []myTypes.CategoryItem
	for _, article := range allArticles {
		formattedPath := formatPath(article.CategoryPath)
		buildCategoryForArticle(&categoryTree, formattedPath, &article)
	}
	return categoryTree
}
