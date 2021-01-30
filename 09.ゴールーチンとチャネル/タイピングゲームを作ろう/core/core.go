package core

import (
	"context"
	"fmt"
	"math"
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
	ctx, timeup := context.WithCancel(context.Background())
	timerChan := time.After(time.Second * time.Duration(c.Env.TimeLimit))
	doneQue := make(chan doneWord, 1024)
	endGameChan := game(ctx, c, doneQue)
	scoreChan := consumeDoneQue(endGameChan, doneQue)
	// configs.wordsからランダムに出題、判定結果はワーカーチャネルに放り込む。制限時間経過で終了.タイムアップまでチャネルに書き込み可能.
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
	fmt.Printf("あなたの成績は・・・%d 問/ 平均%.2f秒　制限時間%d秒 でした！\n\n", len(s.dones), s.averageTime(), c.Env.TimeLimit)
	fmt.Printf("----------------------- 内訳 -----------------------\n")
	for _, dw := range s.dones {
		fmt.Printf("%s: %.2f秒\n", dw.word, math.Round(dw.time.Seconds()*100)/100)
	}
}
