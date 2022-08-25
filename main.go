package main

func main() {
	saito := Judge{"斎藤さん"}
	/*	murata := Murata{Player: Player{"村田さん", 0}}
		yamada := Yamada{Player: Player{"村田さん", 0}}
	*/
	var murata Murata
	var yamada Yamada
	murata.Name = "村田さん"
	yamada.Name = "山田さん"

	var murataTactics Tactics
	murataTactics = randomTactics{}
	murata.setTactics(murataTactics)

	var yamadaTactics Tactics
	yamadaTactics = randomTactics{}
	yamada.setTactics(yamadaTactics)

	saito.startJanken(&murata, &yamada)
}
