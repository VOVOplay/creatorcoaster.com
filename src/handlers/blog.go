package handlers

import (
	"net/http"

	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

type BlogHandler struct {
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

func (h *BlogHandler) HandleLandingBlogsPage(w http.ResponseWriter, r *http.Request) {
	allBlogs := getTestBlogArticles()

	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.BlogLandingPage(allBlogs)
	component.Render(r.Context(), w)
}

func getTestBlogArticles() []myTypes.Article {
	return []myTypes.Article{
		{
			Name:         "blog-one",
			PrettyName:   "Blog One",
			HTML:         "<p>This is blog one</p>",
			CategoryPath: "",
			Info: myTypes.ArticleInfo{
				Date:   "09/03/2026",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "blog-two",
			PrettyName:   "Blog Two",
			HTML:         "<p>This is blog two</p>",
			CategoryPath: "",
			Info: myTypes.ArticleInfo{
				Date:   "09/03/2026",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "blog-three",
			PrettyName:   "Blog Three",
			HTML:         "<p>This is blog three</p>",
			CategoryPath: "",
			Info: myTypes.ArticleInfo{
				Date:   "09/03/2026",
				Author: "VOVOplay",
			},
		},
		{
			Name:         "blog-four",
			PrettyName:   "Blog Four",
			HTML:         "<p>This is blog four</p>",
			CategoryPath: "",
			Info: myTypes.ArticleInfo{
				Date:   "09/03/2026",
				Author: "VOVOplay",
			},
		},
	}
}
