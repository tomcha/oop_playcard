package main

type cyclicTactics struct {
	nowHand int
}

func (t *cyclicTactics) readTactics() int {
	nextHand := t.nowHand
	t.nowHand = t.nowHand%3 + 1
	return nextHand
}
