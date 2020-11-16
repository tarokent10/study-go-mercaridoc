package main

import (
	"time"

	"github.com/tenntenn/greeting/v2"
)

func main() {
	println(greeting.Do(time.Now()))
}
