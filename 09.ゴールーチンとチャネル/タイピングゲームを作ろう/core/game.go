package core

import (
	"context"
	"fmt"
	"math/rand"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/config"
	"study-go--mercaridoc/09.ゴールーチンとチャネル/タイピングゲームを作ろう/io"
	"time"
)

func game(ctx context.Context, c config.Configs, doneQue chan doneWord) chan struct{} {
	finishCh := make(chan struct{}, 0)
	go func() {
		defer close(finishCh)
		var isFinish bool = false

		var isCorrect bool = true
		var q string
		var ans string
		var stime time.Time
		for {
			//出題&結果を突っ込む
			if isCorrect {
				stime = time.Now()
				q = nextWord(c)
			}
			if ans, isFinish = io.ReadStdinWithContext(ctx, q); isFinish {
				break
			} else {
				if q == ans {
					isCorrect = true
				} else {
					isCorrect = false
					fmt.Printf("間違いです。再入力してください。%s:%s\n", q, ans)
					continue
				}
				duration := time.Now().Sub(stime)
				d := &doneWord{
					word: q,
					time: duration,
				}
				doneQue <- *d
			}
		}
		finishCh <- struct{}{}
	}()
	return finishCh
}

func nextWord(c config.Configs) string {
	words := c.Texts.Words
	return words[rand.Intn(len(words)+1)]
}
