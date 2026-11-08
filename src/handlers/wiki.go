package handlers

import (
	"cmp"
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
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

	requestedArticle, err := database.GetWikiArticleByName(requestedArticleName, markdown.GenerateHTMLFromString)
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

	rawCategories, err := database.GetRawWikiCategories()
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	categoryItems := createCategoryItems(rawCategories, allArticles)

	// for caching
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.WikiLayout(requestedArticle, categoryItems)
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

func createCategoryItems(rawCategories []myTypes.RawCategory, allWikiArticles []myTypes.Article) []myTypes.CategoryItem {
	remainingRawCategories := slices.Clone(rawCategories)
	remainingArticles := slices.Clone(allWikiArticles)

	var rootCategories []myTypes.Category
	for _, rawCategory := range rawCategories {
		if rawCategory.ParentCategoryId == myTypes.RootCategoryId { // no parent
			category := buildCategory(rawCategory, &remainingArticles, &remainingRawCategories)
			rootCategories = append(rootCategories, category)
		}
	}

	var rootArticles []myTypes.Article
	for _, article := range allWikiArticles {
		if article.ParentCategoryId == myTypes.RootCategoryId {
			rootArticles = append(rootArticles, article)
		}
	}

	var rootItems []myTypes.CategoryItem
	for i := range rootCategories {
		rootItems = append(rootItems, &rootCategories[i])
	}
	for i := range rootArticles {
		rootItems = append(rootItems, &rootArticles[i])
	}

	rootItems = sortCategoryItems(rootItems)
	return rootItems
}

func buildCategory(rawCategory myTypes.RawCategory, remainingWikiArticles *[]myTypes.Article, remainingRawCategories *[]myTypes.RawCategory) myTypes.Category {
	currentCategory := myTypes.Category{
		Id:        rawCategory.Id,
		Name:      rawCategory.Name,
		SortOrder: rawCategory.SortOrder,
		Elements:  []myTypes.CategoryItem{},
	}

	// add relevant wiki articles
	for i := range slices.Backward(*remainingWikiArticles) {
		wikiArticle := (*remainingWikiArticles)[i]

		if wikiArticle.ParentCategoryId == currentCategory.Id {
			currentCategory.Elements = append(currentCategory.Elements, &wikiArticle)
			*remainingWikiArticles = append((*remainingWikiArticles)[:i], (*remainingWikiArticles)[i+1:]...) // remove the article from the remaining wiki articles
		}
	}

	// render nested categories
	for i := range slices.Backward(*remainingRawCategories) {
		subRawCategory := (*remainingRawCategories)[i]
		if subRawCategory.ParentCategoryId == currentCategory.Id {
			*remainingRawCategories = append((*remainingRawCategories)[:i], (*remainingRawCategories)[i+1:]...) // remove the element from the remaining categories

			subCategory := buildCategory(subRawCategory, remainingWikiArticles, remainingRawCategories)
			currentCategory.Elements = append(currentCategory.Elements, &subCategory)
		}
	}

	currentCategory.Elements = sortCategoryItems(currentCategory.Elements)

	return currentCategory
}

func sortCategoryItems(items []myTypes.CategoryItem) []myTypes.CategoryItem {
	slices.SortFunc(items, func(a myTypes.CategoryItem, b myTypes.CategoryItem) int {
		return cmp.Or(
			cmp.Compare(a.GetSortOrder(), b.GetSortOrder()),
			cmp.Compare(a.GetName(), b.GetName()),
		)
	})
	return items
}
