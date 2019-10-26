package gosrcfmt_test

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/sirkon/message"

	"github.com/sirkon/gosrcfmt"
)

func ExampleAST() {
	fset := token.NewFileSet()
	data, err := ioutil.ReadFile("testdata/main.go")
	if err != nil {
		message.Fatal(err)
	}
	file, err := parser.ParseFile(fset, "testdata/main.go", data, parser.AllErrors|parser.ParseComments)
	if err != nil {
		message.Fatal(err)
	}
	output, err := gosrcfmt.AST(fset, file)
	if err != nil {
		message.Fatal(err)
	}
	_, _ = os.Stdout.Write(output)

	// Output:
	// // +build !windows
	//
	// package pkg
	//
	// func main() {
	// 	if true {
	// 	}
	// }
}
