package myTypes

type Article struct {
	Name       string
	PrettyName string
	HTML       string
	Info       ArticleInfo
}

type ArticleInfo struct {
	Date   string
	Author string
}
