package gneiss

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

func init() {
    spew.Config.Indent = "    "
}

func Execute(filename string, devMode bool) (string, error) {
    filebytes, err := os.ReadFile(filename)
    if err != nil { return "", err }

	l := newLexer(string(filebytes))
	l.lex()

    parser := parser{
    	tokens: l.tokens,
    	pos:    0,
    }

    ast, err := parser.parse()
    if err != nil { return "", err }

    writer := writer{devMode, ast}
    return writer.Write()
}
