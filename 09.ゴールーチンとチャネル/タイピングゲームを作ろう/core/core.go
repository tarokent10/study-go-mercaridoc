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
	return execute(c)
}

func execute(c config.Configs) error {
	// Memo channelはライフサイクルの管理（いつどこでCloseされるか）をしっかりすること！
	// もしかすると発行元がCloseすべき？！
	ctx, timeup := context.WithCancel(context.Background())
	// TODO ゲーム制限時間（秒）をチャネルで
	timerChan := time.After(time.Second * time.Duration(c.Env.TimeLimit))
	// TODO ゲームの結果をワーカーチャネルで(processScoreQueが0になるまでWait)
	doneQue := make(chan doneWord, 1024)
	endGameChan := game(ctx, c, doneQue)
	scoreChan := consumeDoneQue(endGameChan, doneQue)
	// TODO configs.wordsからランダムに出題、判定結果はワーカーチャネルに放り込む。制限時間経過で終了.タイムアップまでチャネルに書き込み可能.
	waitAndTimeup(timerChan, timeup)
	showResult(c, getScore(scoreChan))
	return nil
}
func waitAndTimeup(timerChan <-chan time.Time, timeup context.CancelFunc) {
	select {
	case <-timerChan:
		timeup()
	}
}

func getScore(scoreChan chan score) score {
	defer close(scoreChan)
	return <-scoreChan
}

func showResult(c config.Configs, s score) {
	fmt.Printf("----------------------- Time up -----------------------\n\n")
	fmt.Printf("あなたの回答数は・・・%d 問/ %d秒 でした！\n\n", len(s.dones), c.Env.TimeLimit)
	fmt.Printf("----------------------- 内訳 -----------------------\n")
	for dw := range s.dones {
		// TODO なぜかintに。。
		fmt.Printf("%d", dw)
	}
}
