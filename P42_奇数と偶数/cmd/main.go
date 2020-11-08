package main

import "fmt"

// 奇数と偶数
func main() {
	for i := 1; i <= 100; i++ {
		n := i % 2
		switch n {
		case 0:
			fmt.Printf("%d-偶数\n", i)
		case 1:
			fmt.Printf("%d-奇数\n", i)
		}
	}
}
