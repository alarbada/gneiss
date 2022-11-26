package main

import (
	"flag"
	"fmt"
	"gneiss"
)

var (
    fileF =    flag.String("file", "",    "file to compile")
    devModeF = flag.Bool  ("dev",  false, "dev mode")
)

func main() {
    flag.Parse()

    fmt.Println(*fileF, *devModeF)

    s, err := gneiss.Execute(*fileF, *devModeF)
    if err != nil { panic(err) }

    fmt.Println(s)
}
