package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 前提：ゲーム仕様の解釈
// ゲームの得点：プレイヤーごとに0~100点、ゲーム終了条件は余白部、全ゲーム終了でプレイヤーごとに点数集計
// １ゲームの総得点数:32
const gamePoints int = 32
const totalraounds int = 10

type player struct {
	name   string
	scores map[round]score
}

func (p *player) addScore(r round) {
	if p.getTotalScore() < 100 {
		p.scores[r] = p.scores[r] + 1
	}
}
func (p *player) showScore() int {
	var total int
	for k, v := range p.scores {
		fmt.Printf("%s 第%dラウンド：%d\n", p.name, k, v)
		total += int(v)
	}
	fmt.Printf("%s トータル：%d\n\n", p.name, total)
	return total
}
func (p *player) getTotalScore() int {
	var total int
	for _, s := range p.scores {
		total += int(s)
	}
	return total
}

type round int32
type score int32

type game struct {
	gameround   round
	participant []*player
}

func (g *game) dogame() {
	// game logic ランダムに加点
	fmt.Printf("第%dラウンド開始\n", g.gameround)
	for i := 0; i < gamePoints; i++ {
		t := time.Now().UnixNano()
		rand.Seed(t)
		n := rand.Intn(len(g.participant))
		g.participant[n].addScore(g.gameround)
	}
}

type games struct {
	games []*game
}

func (gs *games) start(rounds int) {
	ps := []*player{{"taro", make(map[round]score)}, {"jiro", make(map[round]score)}, {"saburo", make(map[round]score)}, {"shiro", make(map[round]score)}}
	for i := 1; i <= rounds; i++ {
		g := &game{round(i), ps}
		gs.games = append(gs.games, g)
		g.dogame()
	}
	gs.result(ps)
}
func (gs *games) result(ps []*player) {
	fmt.Printf("end game------------\n\n")
	var winner *player
	var topScore int = 0
	for _, p := range ps {
		s := p.showScore()
		if topScore < s {
			topScore = s
			winner = p
		}
	}
	fmt.Printf("------------------------\n 勝者は%d点で%s!!\n", winner.getTotalScore(), winner.name)
}

func main() {
	gs := games{make([]*game, 0)}
	gs.start(totalraounds)
}
