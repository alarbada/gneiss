package main

import (
	"log"

	"github.com/not-rusty/gneiss"
)

func main() {
    opts := gneiss.ExecuteOptions{
        Dirname:           "./testcases",
        DevMode:           false,
        GenerateTestCases: true,
    }

    err := opts.Exec()
    if err != nil { log.Fatal(err) }

    opts = gneiss.ExecuteOptions{
    	Dirname:           "./testcases",
    	DevMode:           true,
    	GenerateTestCases: true,
    }

    err = opts.Exec()
    if err != nil { log.Fatal(err) }
}
