package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/markdown"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

var InvalidBlogName error = errors.New("Invalid blog name")

type BlogHandler struct {
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

func (h *BlogHandler) HandleBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=60")

	if r.URL.String() == "/blog/" { // blog landing page
		allBlogs := getTestBlogArticles()

		component := views.BlogLandingPage(allBlogs)
		component.Render(r.Context(), w)

		return
	}

	requestedBlogName, err := h.findWhichBlogWasRequested(r)
	if err != nil {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	requestedBlog, err := getTestBlogByName(requestedBlogName)
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
		return splitURL[2], nil
	}
}

func getTestBlogArticles() []myTypes.Article {
	return []myTypes.Article{
		{
			Name:         "blog-one",
			PrettyName:   "Blog One",
			HTML:         markdown.GenerateHTMLFromString(database.LoadStaticMarkdownFile("BLOG-ONE.md")),
			CategoryPath: "",
			Type:         myTypes.BlogArticle,
			Info: myTypes.ArticleInfo{
				Date:          "15/10/2026",
				Author:        "VOVOplay",
				AuthorPFPLink: "https://cdn.discordapp.com/avatars/758322333437394944/30f656e6c904df02ba54a0eebcbedfaa.png?size=1024",
				ReadTime:      "5 minutes",
			},
		},
		{
			Name:         "blog-two",
			PrettyName:   "Blog Two",
			HTML:         "<p>This is blog two</p>",
			CategoryPath: "",
			Type:         myTypes.BlogArticle,
			Info: myTypes.ArticleInfo{
				Date:          "24/12/2026",
				Author:        "Lgx_",
				AuthorPFPLink: "https://cdn.discordapp.com/avatars/778286876619964466/50711b08e645367166da6527f4130c17.png?size=1024",
				ReadTime:      "4 minutes",
			},
		},
		{
			Name:         "blog-three",
			PrettyName:   "Blog Three",
			HTML:         "<p>This is blog three</p>",
			CategoryPath: "",
			Type:         myTypes.BlogArticle,
			Info: myTypes.ArticleInfo{
				Date:          "06/09/2026",
				Author:        "Fona",
				AuthorPFPLink: "https://cdn.discordapp.com/avatars/1319616624780644392/6fdabf0bf5ea3a4bd97fcd62f2740676.png?size=1024",
				ReadTime:      "7 minutes",
			},
		},
		{
			Name:         "blog-four",
			PrettyName:   "Blog Four",
			HTML:         "<p>This is blog four</p>",
			CategoryPath: "",
			Type:         myTypes.BlogArticle,
			Info: myTypes.ArticleInfo{
				Date:          "09/03/2026",
				Author:        "VOVOplay",
				AuthorPFPLink: "https://cdn.discordapp.com/avatars/758322333437394944/30f656e6c904df02ba54a0eebcbedfaa.png?size=1024",
				ReadTime:      "3 minutes",
			},
		},
	}
}

func getTestBlogByName(name string) (myTypes.Article, error) {
	allBlogs := getTestBlogArticles()

	for _, blog := range allBlogs {
		if blog.Name == name {
			return blog, nil
		}
	}

	return myTypes.Article{}, ArticleNotFound
}
