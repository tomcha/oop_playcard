package main

import "fmt"

type askTactics struct{}

func (t askTactics) readTactics() int {
	fmt.Println("何を出しますか？")
	fmt.Println("1: グー")
	fmt.Println("2: チョキ")
	fmt.Println("3: パー")
	fmt.Print("input number >>>")
	var s string
	for {
		fmt.Scan(&s)
		if s == "1" {
			fmt.Println("グーを選びました")
			return stone
		} else if s == "2" {
			fmt.Println("チョキを選びました")
			return scissor
		} else if s == "3" {
			fmt.Println("パーを選びました")
			return paper
		} else {
			fmt.Println("1-3の数字を入力してください")
		}
	}
}
