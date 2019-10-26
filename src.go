package gosrcfmt

import (
	"bytes"
	"go/parser"
	"go/token"
	"io"
)

// Source parse, format and return given source code
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

// SourceWrite writes Source output into dest
func SourceWrite(dest io.Writer, src []byte, filename string) error {
	data, err := Source(src, filename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(dest, bytes.NewBuffer(data)); err != nil {
		return &wrapError{
			msg: "write formatted code",
			err: err,
		}
	}
	return nil
}