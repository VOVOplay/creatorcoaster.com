package myTypes

// Represents both a wiki article and a blog post.
type Article struct {
	Name             string      // The not-pretty name of the article. Used as an identifier in the database and in the URL.
	PrettyName       string      // The pretty name used on the website
	HTML             string      // The generated HTML of the article
	MarkdownSource   string      // The original markdown source of the article
	ParentCategoryId int         // [WIKI] The ID
	SortOrder        int         // The order in which it should be displayed. Any number can be used multiple times, since the primary thing sorting it is still the categories.
	Type             ArticleType // Which type it is (myTypes.ArticleType)
	Info             ArticleInfo
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
var RootCategoryId int = -1

type ArticleType int

const (
	WikiArticle ArticleType = iota
	BlogArticle
)

// The "raw" category, like it is stored in the database. This needs to be built into myTypes.Category before being used
type RawCategory struct {
	Id               int
	Name             string
	SortOrder        int
	ParentCategoryId int
}

// A representation of a category in the wiki sidebar
type Category struct {
	Id        int
	Name      string // The category name
	SortOrder int
	Elements  []CategoryItem // The category's children
}

type CategoryItem interface {
	isCategoryItem()
	GetSortOrder() int
}

func (a *Article) isCategoryItem()  {}
func (c *Category) isCategoryItem() {}

func (a *Article) GetSortOrder() int {
	return a.SortOrder
}

func (c *Category) GetSortOrder() int {
	return c.SortOrder
}
