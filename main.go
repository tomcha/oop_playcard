package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	saito := Judge{"斎藤さん"}
	murata := Player{"村田さん", 0}
	yamada := Player{"山田さん", 0}
	/*	var murata Murata
		var yamada Yamada
		murata.Player.Name = "村田さん"
		murata.Player.WinCount = 0

		yamada.Player.Name = "山田さん"
		yamada.Player.WinCount = 0
	*/
	saito.startJanken(&murata, &yamada)
}
