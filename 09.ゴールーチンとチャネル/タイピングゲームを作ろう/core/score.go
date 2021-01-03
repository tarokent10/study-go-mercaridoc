package core

import (
	"context"
	"time"
)

type (
	score struct {
		dones []doneWord
	}
	doneWord struct {
		word string
		time time.Duration
	}
)

func (s *score) addDoneWord(d doneWord) {
	s.dones = append(s.dones, d)
}
func newDoneWord(word string, time time.Duration) doneWord {
	return *&doneWord{word, time}
}

// 結果処理
// タイムアップまで集計を続ける。タイムアップを検知したらキューに積まれた結果を処理してチャネルを閉じる
func consumeDoneQue(ctx context.Context, doneQue chan doneWord) chan score {
	scoreChan := make(chan score, 0)
	go func(ctx context.Context, doneQue chan doneWord) {
		defer close(doneQue)
		s := &score{dones: make([]doneWord, 0, 0)}
		for {
			var isTimeup bool = false
			select {
			case t := <-doneQue:
				s.addDoneWord(t)
			case <-ctx.Done():
				isTimeup = true
			}
			if isTimeup {
				break
			}
		}
		// タイミングの問題で起こりうる処理こぼし対応（スマートじゃないけど）
		for i := 0; i < len(doneQue); i++ {
			s.addDoneWord(<-doneQue)
		}
		scoreChan <- *s
	}(ctx, doneQue)
	return scoreChan
}
