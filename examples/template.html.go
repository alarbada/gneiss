package main

import (
	"io"

	"github.com/not-rusty/gneiss"
)

func Main(w io.Writer) error {
    i := gneiss.NewInterpreter("./examples/template.html")
    return i.Exec(w)
}
