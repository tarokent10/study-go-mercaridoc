package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type result int

func (r result) value() string {
	var val = ""
	switch r {
	case 0:
		val = "大吉"
	case 1:
		val = "中吉"
	case 2:
		val = "小吉"
	case 3:
		val = "吉"
	case 4:
		val = "凶"
	case 5:
		val = "大凶"
	}
	return val
}

// OmikujiServer is server
type OmikujiServer struct{}

func (o OmikujiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, result(rand.Intn(6)).value())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.Handle("/", &OmikujiServer{})
	http.ListenAndServe(":8080", nil)
}
