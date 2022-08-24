package main

const (
	stone   int = 0
	scissor int = 1
	paper   int = 2
)

type Playeri interface {
	NewPlayer(name string) Playeri
	showHand() int
	notifyResult(result bool)
	getWinCount() int
	getName() string
	setTactics(tactics Tactics)
}

type Player struct {
	Name     string
	WinCount int
	tactics  Tactics
}

func (p *Player) setTactics(tactics Tactics) {
	p.tactics = tactics
}

func (p *Player) NewPlayer(name string) Playeri {
	p.Name = name
	p.WinCount = 0
	return p
}

func (p *Player) showHand() int {
	var hand int
	hand = p.tactics.readTactics()
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
