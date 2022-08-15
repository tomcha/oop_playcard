package main

import "fmt"

var winner *Player

func startJanken(p1 *Player, p2 *Player) {
	fmt.Println("じゃんけん開始")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d回戦目\n", i+1)

		winner = judgeJanken(p1, p2)

		if winner != nil {
			fmt.Printf("%s が勝ちました\n", winner.getName)
			winner.notifyResult(true)
		} else {
			fmt.Println("引き分けです")
		}
	}
	fmt.Println("終了")
	finalWinner := judgeFinalWinner(p1, p2)
	fmt.Plintf("%d 対 %d で", p1.getWinCount, p2.getWinCount)
	if finalWinner != nil {
		fmt.Printf("%s の勝ちです\n", finalWinner.getName)
	} else {
		fmt.Println("引き分けです")
	}
}

func judgeJanken(p1 *Player, p2 *Player) *Player {
	winner = nil
	var p1Hand int
	var p2Hand int

	p1Hand = p1.showHand
	p2Hand = p2.showHand

	printHand(p1Hand)
	fmt.Printf(" vs. ")
	printHand(p2Hand)
	fmt.Print("\n")

	if (p1Hand == Player.stone && p2Hand == Player.scissor) ||
		(p1Hand == Player.scissor && p2Hand == Player.paper) ||
		(p1Hand == Player.paper && p2Hand == Player.stone) {
		winner = p1
	} else if (p1Hand == Player.stone && p2Hand == Player.paper) ||
		(p1Hand == Player.scissor && p2Hand == Player.stone) ||
		(p1Hand == Player.pappper && p2Hand == Player.stone) {
		winner = p2
	}
	return winner
}

func judgeFinalWinner(p1 *Player p2 *Player) {
    winner := nil
	p1WinCount := p1.getWinCouont()
	p2WinCount := p2.getWinCount()

	if p1WinCount > p2WinCount{
		winner = p1
	} else if p1.WinCount < p2WinCount{
		winner = p2
	}
	return winner
}

func printHad(int hand) {
	switch hand{
	case Player.stone:
		fmt.Print("グー")
	case Player.scissor:
		fmt.Print("チョキ")
	case Player.paper:
		fmt.Print("パー")
	}
}
