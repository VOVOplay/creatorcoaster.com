package markdown

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

func GenerateHTMLFromString(sourceString string) string {
	source := []byte(sourceString)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),

		goldmark.WithRendererOptions(
			html.WithHardWraps(),

			renderer.WithNodeRenderers(
				util.Prioritized(&CustomParagraphRenderer{}, 10),
				util.Prioritized(&CustomHeadingRenderer{}, 10),
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

// ---

type CustomParagraphRenderer struct{}

func (r *CustomParagraphRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindParagraph, r.renderParagraph)
}

func (r *CustomParagraphRenderer) renderParagraph(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WriteString(`<p class="md-paragraph">`)
	} else {
		w.WriteString(`</p>`)
	}

	return ast.WalkContinue, nil
}

type CustomHeadingRenderer struct{}

func (r *CustomHeadingRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindHeading, r.renderHeading)
}

func (r *CustomHeadingRenderer) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)

	// shitty name, but basically just h1, h2...
	headingElementPrefix := fmt.Sprintf("h%d", n.Level)

	if entering {
		w.WriteString(`<` + headingElementPrefix + ` class="md-heading">`)
	} else {
		w.WriteString(`</` + headingElementPrefix + `>`)
	}

	return ast.WalkContinue, nil
}
