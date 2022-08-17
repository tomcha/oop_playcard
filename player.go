package main

import "math/rand"

const (
	stone   int = 0
	scissor int = 1
	paper   int = 2
)

type Playeri interface {
}

type Player struct {
	Name     string
	WinCount int
}

func (p *Player) NewPlayer(name string) *Player {
	p.Name = name
	p.WinCount = 0
	return p
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

func (p *Player) notifyResult(result bool) {
	if result == true {
		p.WinCount += 1
	}
}

func (p *Player) getWinCount() int {
	return p.WinCount
}

func (p *Player) getName() string {
	return p.Name
}
