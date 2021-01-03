package core

import (
	"context"
	"fmt"
	"math/rand"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/io"
)

func game(ctx context.Context, c config.Configs, doneQue chan doneWord) chan struct{} {
	finishCh := make(chan struct{}, 0)
	go func() {
		defer close(finishCh)
		var isFinish bool = false

		var isCorrect bool = true
		var q string
		var ans string
		for {
			//出題&結果を突っ込む
			// TODO 誤答時にタイムアップ対応と、平均回答時間の計測
			if isCorrect {
				q = selectWord(c)
			}
			ans = io.ReadStdin(q)
			if q == ans {
				isCorrect = true
			} else {
				isCorrect = false
				fmt.Printf("間違いです。再入力してください。%s:%s\n", q, ans)
				continue
			}
			d := &doneWord{
				word: q,
			}
			doneQue <- *d
			select {
			case <-ctx.Done():
				isFinish = true
			default:
			}
			if isFinish {
				break
			}
		}
		finishCh <- struct{}{}
	}()
	return finishCh
}

func selectWord(c config.Configs) string {
	words := c.Texts.Words
	return words[rand.Intn(len(words)+1)]
}
