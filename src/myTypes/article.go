package myTypes

// Represents both a wiki article and a blog post.
type Article struct {
	Name           string      // The not-pretty name of the article. Used as an identifier in the database and in the URL.
	PrettyName     string      // The pretty name used on the website
	HTML           string      // The generated HTML of the article
	MarkdownSource string      // The original markdown source of the article
	CategoryPath   string      // [WIKI] A unix-like path of the category
	SortOrder      int         // The order in which it should be displayed. Any number can be used multiple times, since the primary thing sorting it is still the categories.
	Type           ArticleType // Which type it is (myTypes.ArticleType)
	Info           ArticleInfo
}

type ArticleInfo struct {
	Date               int64  // The unix time at which the article was published or last updated
	Author             string // [BLOG] The author of the article
	AuthorPFPLink      string // [BLOG] The PFP link of the author
	ReadTime           string // [BLOG] The approximate calculation of the read time
	IsFirstWikiArticle bool   // [WIKI] A flag to denote the first wiki article
}

// Used for making /wiki/ load that, and also to mark it as the first article in article.Info.IsFirstArticle
var FirstArticleName string = "introduction"

type ArticleType int

const (
	WikiArticle ArticleType = iota
	BlogArticle
)

// A representation of a category in the wiki sidebar
type Category struct {
	Name     string         // The category name (pretty)
	Elements []CategoryItem // The category's children
}

type CategoryItem interface {
	isCategoryItem()
}

func (a *Article) isCategoryItem()  {}
func (c *Category) isCategoryItem() {}
