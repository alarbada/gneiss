package main

import (
	"flag"
	"fmt"

	"github.com/not-rusty/gneiss"
)

var (
    fileF =    flag.String("file", "",    "file to compile")
    devModeF = flag.Bool  ("dev",  false, "dev mode")
)

func main() {
    flag.Parse()

    s, err := gneiss.Execute(*fileF, *devModeF)
    if err != nil { panic(err) }

    fmt.Println(s)
}
