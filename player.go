package main

import "math/rand"

const (
	stone   int = 0
	scissor int = 1
	paper   int = 2
)

type Player struct {
	name     string
	winCount int
}

func (p *Player) showHand() int {
	var hand int
	t := rand.Intn(3)
	if t == 0 {
		hand = stone
	} else if t == 1 {
		hand = scissor
	} else {
		hand = paper
	}
	return hand
}

func (p *Player) notifyResult(bool result) {
	if result == true {
		winCount += 1
	}
}

func (p *Player) getWinCount() int {
	return winCount
}
