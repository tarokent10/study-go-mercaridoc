package io

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

// ReadStdin read from stdin
func ReadStdin(msg string) string {
	fmt.Printf("%s\n> ", msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// ReadStdinWithContext read from stdin
// if context is canceled while reading from stdin, return true
func ReadStdinWithContext(ctx context.Context, msg string) (string, bool) {
	var canceled bool
	var ans string
	fmt.Printf("%s\n> ", msg)

	stdinCh := make(chan string)
	go func(ch chan string) {
		defer close(ch)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ch <- scanner.Text()
	}(stdinCh)

	select {
	case txt := <-stdinCh:
		canceled = false
		ans = txt
	case <-ctx.Done():
		canceled = true
		ans = ""
	}
	return ans, canceled
}
