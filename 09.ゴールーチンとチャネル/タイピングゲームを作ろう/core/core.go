package core

import (
	"context"
	"fmt"
	"strings"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/io"
	"time"
)

// GameStart start game
func GameStart(c config.Configs) error {
	ans := io.ReadStdin("***** タイピングゲーム *****\nゲームを開始しますか？  y or n")
	if strings.ToLower(ans) != "y" {
		return nil
	}
	fmt.Printf("タイピングゲーム開始します!!\n\n")
	return execute(context.Background(), c)
}

func execute(ctx context.Context, c config.Configs) error {
	// TODO ゲーム制限時間（秒）をチャネルで
	timeerChan := time.After(time.Second * time.Duration(c.Env.TimeLimit))
	// TODO ゲームの結果をワーカーチャネルで
	processScoreQue := make(chan doneWord)
	scoreChan := processScore(ctx, processScoreQue)
	// TODO configs.wordsからランダムに出題、判定結果はワーカーチャネルに放り込む。制限時間経過で終了.
	select {
	case s := <-scoreChan:
		fmt.Printf("%+v\n", s)
	case <-timeerChan:
		// TOODO
	}
	return nil
}
