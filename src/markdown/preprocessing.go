package markdown

import (
	"strings"
	"unicode/utf8"
)

type ArticleSection struct {
	Name     string
	Contents string
}

// Traverses a string and saves all of the article sections into the database
func SaveArticleSectionsToDatabase(article string) ([]ArticleSection, error) {
	articleSections := getAllArticleSections(article)

	// removing the last \n so it stays clean in the database and in the bot
	for _, articleSection := range articleSections {
		articleSection.Contents = strings.TrimSuffix(articleSection.Contents, "\n")
	}

	// TODO: Save to database

	return articleSections, nil
}

// Removes all of the article section syntax but keeps the contents of that article section in the string
func StripArticleSectionSyntaxFromArticle(article string) string {
	return ""
}

// finds all the article sections and returns them
func getAllArticleSections(article string) []ArticleSection {
	var sections []ArticleSection
	var previousChar rune
	for i := 0; i < len(article); {
		char, width := utf8.DecodeRuneInString(article[i:])

		if char == ':' && previousChar == '\n' {
			articleSection, relativeContinueAt, ok := assembleArticleSection(article[i:])
			if ok {
				sections = append(sections, articleSection)

				nextIndex := i + relativeContinueAt
				if nextIndex > len(article) {
					nextIndex = len(article)
				}

				if nextIndex > 0 {
					previousChar, _ = utf8.DecodeLastRuneInString(article[:nextIndex])
				}

				i = nextIndex
				continue
			}

		}
		previousChar = char
		i += width
	}

	return sections
}

func assembleArticleSection(restOfArticle string) (ArticleSection, int, bool) {
	endOfNameIndex := strings.Index(restOfArticle, "\n")
	if endOfNameIndex == -1 {
		return ArticleSection{}, -1, false
	}

	closingCharacterIndex := -1
	var previousChar rune
	for i, char := range restOfArticle {
		if i > 0 && char == ':' && previousChar == '\n' {
			closingCharacterIndex = i
			break
		}
		previousChar = char
	}

	if closingCharacterIndex == -1 {
		return ArticleSection{}, -1, false
	}

	name := restOfArticle[1:endOfNameIndex]
	name = strings.ReplaceAll(name, " ", "_")

	contents := restOfArticle[endOfNameIndex+1 : closingCharacterIndex]

	relativeContinueAt := closingCharacterIndex + 2

	articleSection := ArticleSection{
		Name:     name,
		Contents: contents,
	}

	return articleSection, relativeContinueAt, true
}
