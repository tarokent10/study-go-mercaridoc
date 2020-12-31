package core

import (
	"context"
	"time"
)

type (
	score struct {
		donns []doneWord
	}
	doneWord struct {
		word string
		time time.Duration
	}
)

func (s *score) addDoneWord(d doneWord) {
	s.donns = append(s.donns, d)
}
func newDoneWord(word string, time time.Duration) doneWord {
	return *&doneWord{word, time}
}

func processScore(ctx context.Context, target <-chan doneWord) <-chan score {
	scoreChan := make(chan score, 0)
	go func(ctx context.Context, target <-chan doneWord) {
		s := &score{}
		for t := range target {
			s.addDoneWord(t)
			select {
			case <-ctx.Done():
				break
			}
		}
		scoreChan <- *s
	}(ctx, target)
	return scoreChan
}
