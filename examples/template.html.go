package main

import "io"

func Main(w io.Writer) {
w.Write([]byte(`
    <div>I'm the first app! omg</div>
    <div>I'm the first app! omg</div>
    <div>I'm the first app! omg</div>
    <div>I'm the first app! omg</div>
`))
}
