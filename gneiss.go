package gneiss

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func init() {
    spew.Config.Indent = "    "
}

func getAllGneissFiles(dir string) (filenames []string, err error) {
    entries, err := os.ReadDir(dir)
    if err != nil { return filenames, err }

    for _, entry := range entries {
        if entry.IsDir() {
            dirname := filepath.Join(dir, entry.Name())
            if err != nil { return filenames, err }

            subfiles, err := getAllGneissFiles(dirname)
            if err != nil { return filenames, err }

            filenames = append(filenames, subfiles...)

            continue
        } else {
            filename := filepath.Join(dir, entry.Name())

            parts := strings.Split(filename, ".")
            if len(parts) == 0 { continue }

            extension := parts[len(parts) - 1]
            if extension != "gneiss" { continue }

            filenames = append(filenames, filename)
        }
    }

    return filenames, nil
}

type ExecuteOptions struct {
    Dirname           string
    DevMode           bool
    GenerateTestCases bool
}

func (x ExecuteOptions) Exec() error {
    files, err := getAllGneissFiles(x.Dirname) 
    if err != nil { return err }

    for _, file := range files {
        filebytes, err := os.ReadFile(file)
        if err != nil { return err }

        l := newLexer(string(filebytes))
        l.lex()

        p := newParser(l.tokens)

        ast, err := p.parse()
        if err != nil { return err }

        writer := writer{x.DevMode, ast}
        if x.GenerateTestCases {
            file := file
            if x.DevMode {
                file += "_interpreted.go"
            } else {
                file += "_compiled.go"
            }

            err = writer.Write(file)
            if err != nil { return err }
        } else {
            err = writer.Write(file + ".go")
            if err != nil { return err }
        }
    }

    return nil
}

func lexAndParse(contents string) (node, error) {
    l := newLexer(contents)
    l.lex()

    p := newParser(l.tokens)
    return p.parse()
}
