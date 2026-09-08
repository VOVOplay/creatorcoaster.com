package markdown

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

func GenerateHTMLFromString(sourceString string) string {
	source := []byte(sourceString)

	md := goldmark.New(
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
			),
		),
	)

	var buf bytes.Buffer
	err := md.Convert(source, &buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf.String())

	return buf.String()
}
