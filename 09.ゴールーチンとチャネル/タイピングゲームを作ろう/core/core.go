package core

import (
	"fmt"
	"strings"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/io"
)

// GameStart start game
func GameStart(c config.Configs) error {
	ans := io.ReadStdin("***** タイピングゲーム *****\nゲームを開始しますか？  y or n")
	if strings.ToLower(ans) != "y" {
		return nil
	}
	fmt.Printf("タイピングゲーム開始します!!\n\n")
	return game(c)
}

func game(c config.Configs) error {
	// TODO ゲーム制限時間（秒）をチャネルで
	// TODO ゲームの結果をワーカーチャネルで
	// TODO configs.wordsからランダムに出題、判定結果はワーカーチャネルに放り込む。制限時間経過で終了.
	return nil
}
