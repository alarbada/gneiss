package gneiss

import (
	"fmt"
	"io"
	"os"
)

type Interpreter struct {
    filecontents string
}

func (x *Interpreter) Read(filename string) {
    bs, err := os.ReadFile(filename)
    if err != nil { panic(err) }

    x.filecontents = string(bs)
}

func (x *Interpreter) Exec(w io.Writer) error {

    l := lexer{
    	tokens:   []token{},
    	contents: x.filecontents,
    	pos:      0,
    }
    l.lex()

    p := parser{
    	tokens: []token{},
    	pos:    0,
    }

    node, err := p.parse()
    if err != nil { return err }
 
    switch n := node.(type) {
    case fileNode:
        for _, child := range n.children {
            switch n := child.(type) {
                case textNode:
                    w.Write([]byte(n.contents))
                case componentNode:
                    // pass
                default:
            }
            
        }
    default:
        return fmt.Errorf("invalid node %v", n)
    }

    return nil
}
