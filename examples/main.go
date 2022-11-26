package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("listening at port :8080")
    http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        err := Main(w)
        if err != nil { fmt.Println(err) }
    }))
}
