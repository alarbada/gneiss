package gneiss

import (
	"fmt"
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

func (x *writer) WriteTmplFile() (string, error) {
    var sb strings.Builder

    switch n := x.ast.(type) {
    case fileNode:
        sb.WriteString(`package main

import "io"

`)

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

func (x *writer) Write() (string, error) {
    if x.devMode {
        return x.WriteTmplFile()
    } else {
        return x.WriteGoFile()
    }
}
