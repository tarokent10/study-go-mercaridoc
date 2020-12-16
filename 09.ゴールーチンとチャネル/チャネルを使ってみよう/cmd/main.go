package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print("> ")
		fmt.Println(<-ch)
	}
}
func input(r io.Reader) <-chan string {
	ch := make(chan string, 0)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		defer close(ch)
	}()
	return ch
}
