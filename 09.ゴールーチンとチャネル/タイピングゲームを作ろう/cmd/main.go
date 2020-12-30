package main

import (
	"log"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/core"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	core.GameStart()
}
