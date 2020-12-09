package main

import (
	"fmt"
	"study-go--mercaridoc/07.エラー処理/エラー処理をしてみよう/errors"
	"study-go--mercaridoc/07.エラー処理/エラー処理をしてみよう/util"
)

func main() {
	do(nil)
	do("文字列")
	do(10)
	do(1.8)
	do(false)
	do(func(s string) {})
	do([]string{"あ", "い"})
	do(errors.NewSampleError("errors"))
}

func do(v interface{}) {
	if s, err := util.ToStringer(v); err != nil {
		fmt.Printf("%T: %s", err, err.Error())
	} else {
		fmt.Printf("%T: %s", s, s.String())
	}
}
