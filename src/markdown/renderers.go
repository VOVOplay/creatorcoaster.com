package markdown

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

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

	// h1, h2...
	headingElementPrefix := fmt.Sprintf("h%d", n.Level)

	if entering {
		w.WriteString(`<` + headingElementPrefix + ` class="md-heading">`)
	} else {
		w.WriteString(`</` + headingElementPrefix + `>`)
	}

	return ast.WalkContinue, nil
}

type CustomBlockQuoteRenderer struct{}

func (r *CustomBlockQuoteRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindBlockquote, r.renderBlockQuote)
}

func (r *CustomBlockQuoteRenderer) renderBlockQuote(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WriteString(`<blockquote class="md-blockquote">`)
	} else {
		w.WriteString(`</blockquote>`)
	}

	return ast.WalkContinue, nil
}

type CustomLargeCodeBlockRenderer struct{}

func (r *CustomLargeCodeBlockRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderLargeCodeBlock)
}

func (r *CustomLargeCodeBlockRenderer) renderLargeCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)

	if entering {
		w.WriteString(`<div class="code-container"><pre class="md-large-code"><code>`)

		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			w.Write(util.EscapeHTML(line.Value(source)))
		}
	} else {
		w.WriteString(`</code></pre><input type="image" src="/static/assets/copy/copy.svg" class="copy-code-button"></input></div>`)
	}

	return ast.WalkContinue, nil
}

type CustomUnorderedListRenderer struct{}

func (r *CustomUnorderedListRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindList, r.renderList)
}

func (r *CustomUnorderedListRenderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.List)

	var startingString string
	var endingString string
	if n.Start == 0 { // (unordered)
		startingString = `<ul class="md-ul">`
		endingString = `</ul>`
	} else {
		startingString = `<ol class="md-ol">`
		endingString = `</ol>`
	}

	if entering {
		w.WriteString(startingString)
	} else {
		w.WriteString(endingString)
	}

	return ast.WalkContinue, nil
}

type CustomSpacerRenderer struct{}

func (r *CustomSpacerRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindSpacer, r.renderSpacer)
}

func (r *CustomSpacerRenderer) renderSpacer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WriteString(`<div class="md-spacer"></div>`)
	}
	return ast.WalkContinue, nil
}

// --- INLINE

type CustomSmallCodeRenderer struct{}

func (r *CustomSmallCodeRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindCodeSpan, r.renderSmallCode)
}

func (r *CustomSmallCodeRenderer) renderSmallCode(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WriteString(`<code class="md-inline-code">`)
	} else {
		w.WriteString(`</code>`)
	}

	return ast.WalkContinue, nil
}
