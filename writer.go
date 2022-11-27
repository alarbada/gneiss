package gneiss

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type writer struct {
    devMode bool
    ast     node
}

func (x *writer) WriteGoFile() (string, error) {
    var sb strings.Builder

    switch n := x.ast.(type) {
    case fileNode:
        _, err := sb.WriteString(`package main

import "io"

`)
        if err != nil { return "", err }

        for _, child := range n.children {
            switch c := child.(type) {
            case componentNode:
                sb.WriteString(`func Main(w io.Writer) {`)
                for _, child := range c.children {
                    switch c := child.(type) {
                    case textNode:
                        sb.WriteString("\nw.Write([]byte(`\n")
                        sb.WriteString(c.contents)
                        sb.WriteString("`))\n")
                    default:
                        return "", fmt.Errorf("invalid node %#v", c)
                    }
                }
                sb.WriteString(`}`)
            case textNode:
                sb.WriteString(c.contents)
            default:
                return "", fmt.Errorf("invalid node %#v", c)
            }
        }


        return sb.String(), nil
    default:
        return "", fmt.Errorf("invalid node %v, expected fileNode", n)
    }
}

func (x *writer) WriteTmplFile() string {
    return `
package main

import (
    "io"
    "github.com/not-rusty/gneiss"
) 

func Main(w io.Writer) {
    i := gneiss.NewInterpreter("./examples/template.html")
    i.Exec(w)
}

`
}

func (x *writer) Write(filename string) error {
    if x.devMode {
        f := x.WriteTmplFile()

        file, err := os.Create(filename)
        if err != nil { return err }
        defer file.Close()

        _, err = file.WriteString(f)
        return err
    } else {
        f, err := x.WriteGoFile()
        if err != nil { return err }

        file, err := os.Create(filename)
        if err != nil { return err }
        defer file.Close()

        _, err = file.WriteString(f)
        return err
    }
}


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

func (x *Interpreter) Exec(w io.Writer) {
    node, err := lexAndParse(x.filecontents)
    if err != nil {
        fmt.Println("interpreting error: ", err)
        return
    }

    err = writeNode(node, w)
    if err != nil {
        fmt.Println("interpreting error: ", err)
    }
}
