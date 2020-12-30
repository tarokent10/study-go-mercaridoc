package io

import (
	"bufio"
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
