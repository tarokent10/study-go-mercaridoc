package core

import (
	"math"
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
func (s *score) averageTime() float64 {
	var sumMin float64
	for _, done := range s.dones {
		sumMin += done.time.Seconds()
	}
	return math.Round((sumMin/float64(len(s.dones)))*100) / 100

}
func newDoneWord(word string, time time.Duration) doneWord {
	return *&doneWord{word, time}
}

// 結果処理
// タイムアップまで集計を続ける。タイムアップを検知したらキューに積まれた結果を処理してチャネルを閉じる
func consumeDoneQue(endGameChan chan struct{}, doneQue chan doneWord) chan score {
	scoreChan := make(chan score, 0)
	go func(endGameChan chan struct{}, doneQue chan doneWord) {
		defer close(doneQue)
		s := &score{dones: make([]doneWord, 0, 0)}
		var isGameEnd bool = false
		for {
			select {
			case t := <-doneQue:
				s.addDoneWord(t)
			case <-endGameChan:
				isGameEnd = true
			}
			if isGameEnd {
				break
			}
		}
		// タイミングの問題で起こりうる処理こぼし対応（スマートじゃないけど）
		for i := 0; i < len(doneQue); i++ {
			s.addDoneWord(<-doneQue)
		}
		scoreChan <- *s
	}(endGameChan, doneQue)
	return scoreChan
}
