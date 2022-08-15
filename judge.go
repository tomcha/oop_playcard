package main

import "fmt"

type Judge struct {
	name string
}

func (j *Judge) startJanken(p1 *Player, p2 *Player) {
	fmt.Println("じゃんけん開始")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d回戦目\n", i+1)

		winner := j.judgeJanken(p1, p2)

		if winner != nil {
			fmt.Printf("%s が勝ちました\n", winner.getName())
			winner.notifyResult(true)
		} else {
			fmt.Println("引き分けです")
		}
	}
	fmt.Println("終了")
	finalWinner := j.judgeFinalWinner(p1, p2)
	fmt.Printf("%d 対 %d で", p1.getWinCount(), p2.getWinCount())
	if finalWinner != nil {
		fmt.Printf("%s の勝ちです\n", finalWinner.getName())
	} else {
		fmt.Println("引き分けです")
	}
}

func (j *Judge) judgeJanken(p1 *Player, p2 *Player) *Player {
	var winner *Player
	var p1Hand int
	var p2Hand int

	p1Hand = p1.showHand()
	p2Hand = p2.showHand()

	j.printHand(p1Hand)
	fmt.Printf(" vs. ")
	j.printHand(p2Hand)
	fmt.Print("\n")

	if (p1Hand == stone && p2Hand == scissor) ||
		(p1Hand == scissor && p2Hand == paper) ||
		(p1Hand == paper && p2Hand == stone) {
		winner = p1
	} else if (p1Hand == stone && p2Hand == paper) ||
		(p1Hand == scissor && p2Hand == stone) ||
		(p1Hand == paper && p2Hand == stone) {
		winner = p2
	}
	return winner
}

func (j *Judge) judgeFinalWinner(p1 *Player, p2 *Player) *Player {
	var winner *Player
	p1WinCount := p1.getWinCount()
	p2WinCount := p2.getWinCount()

	if p1WinCount > p2WinCount {
		winner = p1
	} else if p1.WinCount < p2WinCount {
		winner = p2
	}
	return winner
}

func (j *Judge) printHand(hand int) {
	switch hand {
	case stone:
		fmt.Print("グー")
	case scissor:
		fmt.Print("チョキ")
	case paper:
		fmt.Print("パー")
	}
}
