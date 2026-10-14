package markdown

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var KindVariable = ast.NewNodeKind("Variable")

type Variable struct {
	ast.BaseInline
	variableName []byte
}

func NewVariable(name []byte) *Variable {
	return &Variable{
		variableName: name,
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
	fmt.Println(end)
	if end < 0 {
		return nil // returning nil falls back to normal text
	}

	variableValue := line[1:end]

	fmt.Println(string(variableValue))

	block.Advance(end + 1)

	return NewVariable(variableValue)
}
