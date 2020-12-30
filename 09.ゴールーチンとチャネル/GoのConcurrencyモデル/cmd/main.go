package main

import (
	"fmt"
	"time"
)

// （たぶん）スタンダード？なGoによる並行処理パターン
var done = make(chan bool) // 並行処理完了のお知らせ用.よく使われるパターンぽい
var msgs = make(chan int)

func producer() {
	for i := 0; i < 500; i++ {
		time.Sleep(time.Millisecond * 20)
		msgs <- i
	}
	done <- true
}

func consumer() {
	for {
		msg := <-msgs
		fmt.Println(msg)
	}
}

func main() {
	go producer()
	go consumer()
	<-done // 1:1 or 1:nで子チャンネルから受ける使い方が多い？
	println("done!!")
}
