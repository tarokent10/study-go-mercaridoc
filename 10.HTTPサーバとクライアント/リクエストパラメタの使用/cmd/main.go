package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if v := r.FormValue("p"); v == "Gopher" {
			fmt.Fprintln(w, "Goppherさんの運勢は「大吉」です！")
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
