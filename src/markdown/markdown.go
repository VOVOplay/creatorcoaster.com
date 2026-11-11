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

var UGCPolicy bluemonday.Policy = *bluemonday.UGCPolicy()

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

	html := UGCPolicy.Sanitize(buf.String())

	return html
}
