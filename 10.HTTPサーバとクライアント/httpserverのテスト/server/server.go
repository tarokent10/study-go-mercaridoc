package server

import (
	"fmt"
	"net/http"
)

// SampleHandler is handler
func SampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,World")
}
