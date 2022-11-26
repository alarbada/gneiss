package gneiss

import (
	"io"
	"os"
)

type Interpreter struct {
    filecontents string
}

func NewInterpreter(filename string) Interpreter {
    bs, err := os.ReadFile(filename)
    if err != nil { panic(err) }

    return Interpreter{
    	filecontents: string(bs),
    }
}

func writeNode(n node, w io.Writer) error {
    switch n := n.(type) {
    case fileNode:
        for _, child := range n.children {
            err := writeNode(child, w)
            if err != nil { return err }
        }
    case componentNode:
        for _, child := range n.children {
            err := writeNode(child, w)
            if err != nil { return err }
        }
    case textNode:
        _, err := w.Write([]byte(n.contents))
        if err != nil { return err }
    }

    return nil
}

func (x *Interpreter) Exec(w io.Writer) error {
    node, err := lexAndParse(x.filecontents)
    if err != nil { return err }

    return writeNode(node, w)
}
