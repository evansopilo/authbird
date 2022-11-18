package main

import (
	"fmt"
	"net/http"
)

// Wildcard route traps all /api/* requests and response with URL Path
func (app App) Wildcard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path)
}
