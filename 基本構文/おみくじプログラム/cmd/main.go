package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type omikujiResult struct {
	dice    int
	omikuji string
}

func main() {
	fmt.Println("おみくじを引きますか？ y or n")
	for {
		var s string
		fmt.Print("> ")
		if _, err := fmt.Scan(&s); err == nil {
			if strings.ToLower(s) == "y" {
				fmt.Printf("おみくじの結果は")
				for i := 0; i < 3; i++ {
					time.Sleep(time.Millisecond * 250)
					fmt.Print("・・・")
					time.Sleep(time.Millisecond * 250)
				}
				r := omikuji()
				fmt.Printf("ダイス%d 『%s』でした!!\n", r.dice, r.omikuji)
				fmt.Println("もう一度おみくじを引きますか？ y or n")
			} else {
				fmt.Println("おしまい")
				break
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

func omikuji() omikujiResult {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6) + 1

	var r omikujiResult
	switch s {
	case 6:
		r = omikujiResult{s, "大吉"}
	case 5, 4:
		r = omikujiResult{s, "中吉"}
	case 3, 2:
		r = omikujiResult{s, "吉"}
	case 1:
		r = omikujiResult{s, "凶"}
	}
	return r
}
