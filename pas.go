package pas

import (
	"github.com/akm/pas/ast"
	"github.com/akm/pas/parser"
)

func ParseString(code string) (*ast.File, error) {
	return parser.ParseString(code)
}
