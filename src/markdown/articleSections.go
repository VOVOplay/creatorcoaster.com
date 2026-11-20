package markdown

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
)

var ErrSectionsMustBeInWikiArticles error = errors.New("Article sections can only be in wiki articles.")

type articleSectionLineInfo struct {
	openingLine line
	closingLine line
}

type line struct {
	firstIndex int
	lastIndex  int
}

// Traverses a string and saves all of the article sections into the database
func ProcessAndSaveArticleSections(article myTypes.Article) ([]myTypes.ArticleSection, error) {
	if article.Type != myTypes.WikiArticle {
		return []myTypes.ArticleSection{}, ErrSectionsMustBeInWikiArticles
	}

	articleSections, _ := getAllArticleSections(article.MarkdownSource)

	for i, articleSection := range articleSections {
		articleSections[i].Contents = strings.TrimSuffix(articleSections[i].Contents, "\n") // removing the last \n so it stays clean in the database and in the bot
		articleSection.Article = article

		database.SaveArticleSection(articleSection)
	}

	return articleSections, nil
}

// Removes all of the article section syntax but keeps the contents of that article section in the string
func StripArticleSectionSyntaxFromArticle(article string) string {
	_, sectionsLineInfo := getAllArticleSections(article)

	var builder strings.Builder

	for i, char := range article {
		if isPartOfSyntax(i, sectionsLineInfo) {
			continue
		}
		builder.WriteRune(char)
	}

	return builder.String()
}

func isPartOfSyntax(i int, sectionsLineInfo []articleSectionLineInfo) bool {
	for _, info := range sectionsLineInfo {
		if info.isMeantForRemoval(i) {
			return true
		}
	}
	return false
}

func (l line) contains(index int) bool {
	return index >= l.firstIndex && index <= l.lastIndex
}

func (info articleSectionLineInfo) isMeantForRemoval(index int) bool {
	return info.openingLine.contains(index) || info.closingLine.contains(index)
}

// finds all the article sections and returns them
func getAllArticleSections(article string) ([]myTypes.ArticleSection, []articleSectionLineInfo) {
	var sections []myTypes.ArticleSection
	var sectionsLineInfo []articleSectionLineInfo
	var previousChar rune

	for i := 0; i < len(article); {
		char, width := utf8.DecodeRuneInString(article[i:])

		if char == ':' && (i == 0 || previousChar == '\n') {
			articleSection, endOfNameIndex, relativeClosingIndex, hasTrailingNewline, ok := assembleArticleSection(article[i:])
			if ok {
				sections = append(sections, articleSection)

				closingIndex := i + relativeClosingIndex
				nextIndex := closingIndex + 1
				if hasTrailingNewline {
					nextIndex += 1
				}

				sectionLineInfo := articleSectionLineInfo{
					openingLine: line{
						firstIndex: i,
						lastIndex:  i + endOfNameIndex,
					},
					closingLine: line{
						firstIndex: closingIndex,
						lastIndex:  closingIndex + 1,
					},
				}

				sectionsLineInfo = append(sectionsLineInfo, sectionLineInfo)

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

	return sections, sectionsLineInfo
}

func assembleArticleSection(restOfArticle string) (myTypes.ArticleSection, int, int, bool, bool) {
	endOfNameIndex := strings.Index(restOfArticle, "\n")
	if endOfNameIndex == -1 {
		return myTypes.ArticleSection{}, -1, -1, false, false
	}

	closingCharIndex := -1 // the colon on the last line
	var previousChar rune
	for i, char := range restOfArticle {
		if i > 0 && char == ':' && previousChar == '\n' {
			closingCharIndex = i
			break
		}
		previousChar = char
	}

	if closingCharIndex == -1 {
		return myTypes.ArticleSection{}, -1, -1, false, false
	}

	name := restOfArticle[1:endOfNameIndex]
	name = strings.ReplaceAll(name, " ", "_")
	if len(name) < 1 {
		return myTypes.ArticleSection{}, -1, -1, false, false
	}

	contents := restOfArticle[endOfNameIndex+1 : closingCharIndex]

	articleSection := myTypes.ArticleSection{
		Name:     name,
		Contents: contents,
	}

	hasTrailingNewline := closingCharIndex+1 < len(restOfArticle) && restOfArticle[closingCharIndex+1] == '\n'

	return articleSection, endOfNameIndex, closingCharIndex, hasTrailingNewline, true
}
