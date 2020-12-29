package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	go func(t *time.Ticker) {
		time.Sleep(time.Second * 10)
		t.Stop()
	}(t)
	for {
		select {
		case <-t.C:
			process()
		case <-time.After(time.Second * 10):
			fmt.Println("timeout")
			break
		}
	}

}
func process() {
	fmt.Println("定期処理")
}
