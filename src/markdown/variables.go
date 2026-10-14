package markdown

import (
	"bytes"
	"fmt"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var KindVariable = ast.NewNodeKind("Variable")

type Variable struct {
	ast.BaseInline
	Name []byte
}

func NewVariable(name []byte) *Variable {
	return &Variable{
		Name: name,
	}
}

func (n *Variable) Kind() ast.NodeKind {
	return KindVariable
}

func (n *Variable) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// ---

type VariableExtension struct{}

func (e *VariableExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(&VariableParser{}, 500),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&CustomVariableRenderer{}, 500),
		),
	)
}

// ---

type VariableParser struct{}

func (p *VariableParser) Trigger() []byte {
	return []byte{'{'}
}

func (p *VariableParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	// line is everything from { to the end of the line.
	line, _ := block.PeekLine()

	end := bytes.IndexByte(line, '}')
	if end < 0 {
		return nil // returning nil falls back to normal text
	}

	variableName := line[1:end]

	block.Advance(end + 1)

	return NewVariable(variableName)
}

// ---

type CustomVariableRenderer struct{}

func (r *CustomVariableRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindVariable, r.renderVariable)
}

func (r *CustomVariableRenderer) renderVariable(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*Variable)
		variableValue, err := database.GetVariableByName(string(n.Name))
		if err != nil {
			fmt.Fprintf(w, "{%v}", string(n.Name)) // "replicating" the original look, there might be a better way to do that but fuck that
			return ast.WalkContinue, nil
		}

		w.WriteString(variableValue)
	}

	return ast.WalkContinue, nil
}
