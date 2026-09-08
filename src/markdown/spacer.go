package markdown

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
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

// ---

type SpacerTransformer struct{}

func (t *SpacerTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	var targets []ast.Node

	// find all empty lines (is a block element and the previous line is blank)
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && n.Type() == ast.TypeBlock && n.PreviousSibling() != nil {
			if n.HasBlankPreviousLines() {
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
