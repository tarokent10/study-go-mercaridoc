package main

import (
	"context"
	"fmt"
)

type (
	ctxkey string // コンテキストValueのキーは自由にSetされるとカオスなので公開しない
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	bc := context.Background()
	vctx := context.WithValue(bc, ctxkey("key"), "value☆")
	ctx, cancel := context.WithCancel(vctx)
	defer cancel()
	for n := range gen(ctx) { // channelもrangeできる
		fmt.Printf("%d\n", n)
		if n == 5 {
			break
		}
	}
	fmt.Printf("value is %s\n", vctx.Value(ctxkey("key")))
	fmt.Println("done!")
}
