package markdown

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var KindSpacer = ast.NewNodeKind("Spacer")

type Spacer struct {
	ast.BaseBlock
}

func NewSpacer() *Spacer {
	return &Spacer{}
}

func (n *Spacer) Kind() ast.NodeKind {
	return KindSpacer
}

func (n *Spacer) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

type SpacerTransformer struct{}

func (t *SpacerTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	var targets []ast.Node
	source := reader.Source()

	// find all empty lines (is a block element and the previous line is blank)
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && n.Type() == ast.TypeBlock && n.PreviousSibling() != nil {
			if n.HasBlankPreviousLines() {
				previousSibling := n.PreviousSibling()

				// goldmark adds a new line after a fenced code block
				if previousSibling.Kind() == ast.KindFencedCodeBlock {
					prevLines := previousSibling.Lines()
					currLines := n.Lines()

					if prevLines.Len() > 0 && currLines.Len() > 0 {
						endOfPrev := prevLines.At(prevLines.Len() - 1).Stop
						startOfCurr := currLines.At(0).Start

						between := source[endOfPrev:startOfCurr]

						if !hasActualBlankLine(between) {
							return ast.WalkContinue, nil
						}
					}
				}

				targets = append(targets, n)
			}
		}
		return ast.WalkContinue, nil
	})

	// add the spacers to the AST
	for _, n := range targets {
		spacer := NewSpacer()
		n.Parent().InsertBefore(n.Parent(), n, spacer)
	}
}

func hasActualBlankLine(b []byte) bool {
	newlines := 0
	for _, char := range b {
		if char == '\n' {
			newlines++
			if newlines >= 2 {
				return true
			}
		}
	}
	return false
}

// ---

// This extension parses and renders tiny text syntax.
// A tiny text has to start on a new line and follow this: '-# '.
// Everything to the end of the line will be made small.
type TinyTextExtension struct{}

func (e *TinyTextExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(&TinyTextParser{}, 500),
		),
	)

	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&TinyTextRenderer{}, 500),
		),
	)
}

var KindTinyText = ast.NewNodeKind("Tiny Text")

type TinyText struct {
	ast.BaseBlock
}

func NewTinyText() *TinyText {
	return &TinyText{}
}

func (n *TinyText) Kind() ast.NodeKind {
	return KindTinyText
}

func (n *TinyText) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

type TinyTextParser struct{}

func (p *TinyTextParser) Trigger() []byte {
	return []byte{'-'}
}

func (p *TinyTextParser) CanInterruptParagraph() bool { return true }

// Not sure about this yet, gotta test it first
func (p *TinyTextParser) CanAcceptIndentedLine() bool { return true }

//---

func (p *TinyTextParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()

	if len(line) < 3 || line[1] != '#' || line[2] != ' ' {
		return nil, parser.NoChildren
	}

	segment.Start += 3

	reader.Advance(segment.Len())

	node := NewTinyText()
	node.Lines().Append(segment)

	return node, parser.NoChildren
}

func (p *TinyTextParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (p *TinyTextParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {}

// ---

type TinyTextRenderer struct{}

func (r *TinyTextRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindTinyText, r.renderTinyText)
}

func (r *TinyTextRenderer) renderTinyText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		fmt.Fprint(w, `<span class="md-tiny-text">`)
	} else {
		fmt.Fprintf(w, `</span>`)
	}
	return ast.WalkContinue, nil
}
