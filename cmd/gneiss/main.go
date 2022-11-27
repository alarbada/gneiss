package main

import (
	"flag"
	"log"

	"github.com/not-rusty/gneiss"
)

var (
    dirF      = flag.String("dir",      ".",   "dir to compile gneiss templates")
    devModeF  = flag.Bool  ("dev",      false, "dev mode")
)

func main() {
    flag.Parse()

    opts := gneiss.ExecuteOptions{
    	Dirname:           *dirF,
    	DevMode:           *devModeF,
    	GenerateTestCases: false,
    }

    err := opts.Exec()
    if err != nil { log.Fatal(err) }
}
