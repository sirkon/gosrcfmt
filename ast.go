package gosrcfmt

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

// AST format of given AST
func AST(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var dest bytes.Buffer
	printConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
	if err := printConfig.Fprint(&dest, fset, file); err != nil {
		return nil, err
	}
	return dest.Bytes(), nil
}

// ASTWrite writes AST output into dest
func ASTWrite(dest io.Writer, fset *token.FileSet, file *ast.File) error {
	res, err := AST(fset, file)
	if err != nil {
		return err
	}
	if _, err := io.Copy(dest, bytes.NewBuffer(res)); err != nil {
		return &wrapError{
			msg: "write formatted code",
			err: err,
		}
	}
	return nil
}