package main

import (
	"html/template"
	"net/http"
)

var (
	tpml *template.Template
)

func main() {

	tpml = template.Must(template.New(("Unsei")).Parse("<html><body>Gopherさんの運勢は「<b>大吉</b>」です</body></html>"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpml.Execute(w, r.FormValue("hogehoge")) // 埋め込みもできる
	})
	http.ListenAndServe(":8080", nil)
}
