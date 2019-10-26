package gosrcfmt

import (
	"go/parser"
	"go/token"
)

// Source форматирование данного src
func Source(src []byte, filename string) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return nil, &wrapError{
			msg: "parse input data as Go source code",
			err: err,
		}
	}
	res, err := AST(fset, file)
	if err != nil {
		return nil, &wrapError{
			msg: "format AST",
			err: err,
		}
	}

	return res, nil
}
