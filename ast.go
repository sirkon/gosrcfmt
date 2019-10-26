package gosrcfmt

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

// AST форматирование кода заданного в AST
func AST(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var dest bytes.Buffer
	printConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
	if err := printConfig.Fprint(&dest, fset, file); err != nil {
		return nil, err
	}
	return dest.Bytes(), nil
}
