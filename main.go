package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	saito := Judge{"斎藤さん"}
	murata := Murata{Player: Player{"村田さん", 0}}
	yamada := Yamada{Player: Player{"村田さん", 0}}

	saito.startJanken(&murata, &yamada)
}
