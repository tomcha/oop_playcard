package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	stone   int = 0
	scissor int = 1
	paper   int = 2
)

func main() {

	var player1Hand int
	var player2Hand int
	var j int

	var p1WinCount int
	var p2WinCount int

	fmt.Println("じゃんけん開始")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d回戦\n", i)
		rand.Seed(time.Now().UnixNano())
		j = rand.Intn(3)
		if j == 0 {
			player1Hand = stone
			fmt.Println("P1 グー")
		} else if j == 1 {
			player1Hand = scissor
			fmt.Println("P1 チョキ")
		} else {
			player1Hand = paper
			fmt.Println("P1 パー")
		}

		j = rand.Intn(3)
		if j == 0 {
			player2Hand = stone
			fmt.Println("P2 グー")
		} else if j == 1 {
			player2Hand = scissor
			fmt.Println("P2 チョキ")
		} else {
			player2Hand = paper
			fmt.Println("P2 パー")
		}

		if (player1Hand == 0 && player2Hand == 1) || (player1Hand == 1 && player2Hand == 2) || (player1Hand == 2 && player2Hand == 0) {
			fmt.Println("P1の勝ち")
			p1WinCount += 1
		} else if player1Hand != player2Hand {
			fmt.Println("P2の勝ち")
			p2WinCount += 1
		} else {
			fmt.Println("引き分け")
		}
		fmt.Println()
	}
	fmt.Println("じゃんけん終了\n")
	if p1WinCount > p2WinCount {
		fmt.Printf("%d vs %d でP1の勝ち\n", p1WinCount, p2WinCount)
	} else if p1WinCount < p2WinCount {
		fmt.Printf("%d vs %d でP2の勝ち\n", p1WinCount, p2WinCount)
	} else {
		fmt.Printf("%d vs %d で引き分け\n", p1WinCount, p2WinCount)
	}
}
