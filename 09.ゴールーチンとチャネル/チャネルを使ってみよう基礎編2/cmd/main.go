package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	// var ch2 chan string // ゼロ値なのでnil
	go func() { ch1 <- 100 }()
	go func() { ch2 <- "hi" }()

	// switchではない
	for i := 0; i < 2; i++ {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
		case v2 := <-ch2:
			fmt.Println(v2)
		case <-time.After(5 * time.Second):
			// time.After(): 5秒立ったら現在時刻が送られてくるチャネルを返す
			fmt.Println("time out!!")
		}
	}
}
