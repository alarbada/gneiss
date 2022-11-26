package testdata

import "io"

func Main(w io.Writer) {
w.Write(`
    <h1>what</h1>
    <div>x is something</div>
`)
}
