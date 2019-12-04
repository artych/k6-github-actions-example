package main

import (
	"net/http"

	"github.com/artych/k6-github-actions-example/hello"
)

func main() {
	http.HandleFunc("/", hello.Handler)
	http.ListenAndServe(":8080", nil)
}
