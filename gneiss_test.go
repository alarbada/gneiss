package gneiss_test

import (
	"os"
	"path/filepath"
	"testing"

	"gneiss"
)

func TestItWorks(t *testing.T) {
    paths, err := filepath.Glob(filepath.Join("testdata", "*.html"))
    if err != nil { t.Fatal(err) }

    for _, path := range paths {
        _, filename := filepath.Split(path)
        testname := filename[:len(filename)-len(filepath.Ext(path))]

        t.Run(testname, func(t *testing.T) {
            got, err := gneiss.Execute(path, false)
            if err != nil { t.Fatal("execution error:", err) }

            goFilePath := filepath.Join("testdata", testname+".go")

            want, err := os.ReadFile(goFilePath)
            if err != nil { t.Fatal("error reading golden file:", err) }

            if got != string(want) {
                t.Errorf("\n==== got:\n%s\n==== want:\n%s\n", got, want)
            }
        })
    }
}
