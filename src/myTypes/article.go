package myTypes

type Article struct {
	Name         string
	PrettyName   string
	HTML         string
	CategoryPath string
	Type         ArticleType
	Info         ArticleInfo
}

type ArticleInfo struct {
	Date               string
	Author             string
	AuthorPFPLink      string
	ReadTime           string
	IsFirstWikiArticle bool
}

// Used for making /wiki/ load that, and also to mark it as the first article in article.Info.IsFirstArticle
var FirstArticleName string = "introduction"

type ArticleType int

const (
	WikiArticle ArticleType = iota
	BlogArticle
)

type Category struct {
	Name     string
	Elements []CategoryItem
}

type CategoryItem interface {
	isCategoryItem()
}

func (a *Article) isCategoryItem()  {}
func (c *Category) isCategoryItem() {}
