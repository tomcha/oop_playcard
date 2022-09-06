package main

type cyclicTactics struct {
	nowHand int
}

func (t *cyclicTactics) readTactics() int {
	nextHand := t.nowHand
	if t.nowHand == stone {
		t.nowHand = scissor
	} else if t.nowHand == scissor {
		t.nowHand = paper
	} else {
		t.nowHand = stone
	}
	return nextHand
}
