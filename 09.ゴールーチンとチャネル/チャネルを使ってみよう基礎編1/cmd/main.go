package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 100
	}()
	go func() {
		fmt.Println("待ちます")
		v := <-ch
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)

}
