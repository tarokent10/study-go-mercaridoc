package main

import (
	"bytes"
	"fmt"
	"log"
	"study-go--mercaridoc/07.エラー処理/エラー処理をまとめる/util"
)

func main() {
	bf := bytes.NewBufferString("ハローworld!")
	s := util.NewRuneScanner(bf)
	for s.Scan() {
		rune := s.Rune()
		fmt.Printf("%s", string(rune))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err.Error())
	}
}
