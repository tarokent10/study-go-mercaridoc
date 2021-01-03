package main

import (
	"log"
	"math/rand"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/core"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := core.GameStart(c); err != nil {
		log.Fatal(err)
	}
}
