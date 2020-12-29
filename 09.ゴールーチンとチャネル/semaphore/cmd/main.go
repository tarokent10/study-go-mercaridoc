package main

import (
	"fmt"
	"time"
)

// go routineによるセマフォの実装パターンを試す
// 最大入室５人までチャットルームでログイン中ユーザと総数を出す
const (
	maxConcurrency int = 5
	jobsize        int = 500
)

var (
	doneIds     chan string   = make(chan string, jobsize)
	nowloginIDs chan string   = make(chan string, jobsize)
	sem         chan struct{} = make(chan struct{}, maxConcurrency)
)

func main() {
	// 20桁ID採番
	ids := func() []string {
		s := jobsize
		rs := make([]string, 0, s)
		for i := 1; i <= s; i++ {
			rs = append(rs, fmt.Sprintf("%020d", i))
		}
		return rs
	}()

	// 同時実行数制御のもとログイン実行
	for _, id := range ids {
		go login(id)
	}

	// 終了待ち
	for i := 0; i < cap(doneIds); i++ {
		select {
		case <-doneIds:
		}
	}
	println("done!!!")
}

func login(id string) {
	sem <- struct{}{}
	// この中の処理はmaxConcurrencyの数しか同時実行できない.
	time.Sleep(time.Millisecond * 1000)
	nowloginIDs <- id
	fmt.Printf("logined.id=%s counter=%d\n", id, len(nowloginIDs))
	<-nowloginIDs
	<-sem
	doneIds <- id
}

func enque(ss []string, s string) []string {
	ret := append(ss, s)
	return ret
}
func deque(ss []string) []string {
	if len(ss) == 0 {
		return ss
	}
	return ss[1:]
}
