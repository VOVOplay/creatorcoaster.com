package myTypes

type Article struct {
	Name         string
	PrettyName   string
	HTML         string
	CategoryPath string
	Info         ArticleInfo
}

type ArticleInfo struct {
	Date   string
	Author string
}

type Category struct {
	Name     string
	Elements []CategoryItem
}

type CategoryItem interface {
	isCategoryItem()
}

func (a *Article) isCategoryItem()  {}
func (c *Category) isCategoryItem() {}

// Articles:
// []Article
// Input: /something/perks
// Resolves path on the fly? Or struct for

/*

Category {
	Article
	Article
	Article
	Category {
		Article
		Article
	}
	Article
	Category {
		Article
		Article
		Article
	}
}

Category {
	Article
	Article
}

*/
