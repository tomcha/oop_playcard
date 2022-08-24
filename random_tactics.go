package main

import (
	"math/rand"
	"time"
)

type randomTactics struct {
}

func (t randomTactics) readTactics() int {
	rand.Seed(time.Now().UnixNano())
	hand := rand.Intn(3)
	if hand == 0 {
		return stone
	} else if hand == 1 {
		return scissor
	} else {
		return paper
	}
}
