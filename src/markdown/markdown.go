package markdown

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

func configurePolicy() bluemonday.Policy {
	policy := *bluemonday.NewPolicy()

	policy.AllowElements(
		"p", "div", "span", "br", "hr",
		"h1", "h2", "h3", "h4", "h5", "h6",
		"strong", "em", "b", "i", "u", "s", "del", "mark", "sub", "sup",
		"ul", "ol", "li", "blockquote", "code", "pre",
		"a", "img",
	)

	policy.AllowStandardURLs()
	policy.AllowAttrs("href", "title", "target").OnElements("a")

	policy.AllowStyling()

	return policy
}

// Generates HTML from a string source.
// This uses all of the custom styling.
func GenerateHTMLFromString(sourceString string) string {
	source := []byte(sourceString)

	md := goldmark.New(
		goldmark.WithExtensions(
			&VariableExtension{},
			&TinyTextExtension{},
		),

		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&SpacerTransformer{}, 10),
			),
		),

		goldmark.WithRendererOptions(
			html.WithHardWraps(),

			renderer.WithNodeRenderers(
				util.Prioritized(&CustomParagraphRenderer{}, 10),
				util.Prioritized(&CustomHeadingRenderer{}, 10),
				util.Prioritized(&CustomBlockQuoteRenderer{}, 10),
				util.Prioritized(&CustomLargeCodeBlockRenderer{}, 10),
				util.Prioritized(&CustomUnorderedListRenderer{}, 10),
				util.Prioritized(&CustomSmallCodeRenderer{}, 10),
				util.Prioritized(&CustomSpacerRenderer{}, 10),
				util.Prioritized(&CustomHorizontalRuleRenderer{}, 10),
			),
		),
	)

	var buf bytes.Buffer
	err := md.Convert(source, &buf)
	if err != nil {
		return `<p class="md-paragraph">Could not generate HTML.</p>`
	}

	policy := configurePolicy()
	html := policy.Sanitize(buf.String())

	return html
}
