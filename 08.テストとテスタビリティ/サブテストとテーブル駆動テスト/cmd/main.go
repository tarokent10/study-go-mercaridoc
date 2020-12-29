package main

import "fmt"

func main() {
	h := Hex(80)
	println(h.string())
}

// Hex is hex
type Hex int

func (h Hex) string() string {
	return fmt.Sprintf("%x", int(h))
}
