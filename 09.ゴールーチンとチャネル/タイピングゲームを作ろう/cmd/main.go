package main

import (
	"log"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/core"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := core.GameStart(c); err != nil {
		log.Fatal(err)
	}
}
