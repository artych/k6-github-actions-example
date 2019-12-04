package hello

import (
	"fmt"
	"net/http"
)

// Greet Greets GitHub Actions
func Greet(path string) string {
	return fmt.Sprintf("Hello GitHub Actions, %s!", path)
}

// Handler - basic http hanler
func Handler(w http.ResponseWriter, r *http.Request) {
	greet := Greet(r.URL.Path[1:])
	fmt.Fprintf(w, greet)
}
