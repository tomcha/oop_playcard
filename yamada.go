package main

type Yamada struct {
	Player
}

func (y *Yamada) showHand() int {
	return paper
}
