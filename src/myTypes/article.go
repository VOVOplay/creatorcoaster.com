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

// Articles:
// []Article
// Input: /something/perks
// Resolves path on the fly? Or struct for
