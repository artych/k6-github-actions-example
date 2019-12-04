package main

import (
	"fmt"

	"github.com/artych/k6-github-actions-example/hello"
)

func main() {
	fmt.Println(hello.Greet())
}
