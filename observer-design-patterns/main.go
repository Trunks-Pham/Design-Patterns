package main

func main() {
	match := NewMatch()

	player1 := &Player{Name: "- Thảo pick Yasuo"}
	player2 := &Player{Name: "- Xin Yasuo nhaa"}
	player3 := &Player{Name: "- Thua tại Yasuo"}

	match.Attach(player1)
	match.Attach(player2)
	match.Attach(player3)

	yasuo := Champion{Name: "Yasuo", Role: "Mid"}
	jinx := Champion{Name: "Jinx", Role: "ADC"}
	thresh := Champion{Name: "Thresh", Role: "Support"}

	match.SelectChampion("- Thảo pick Yasuo", yasuo)
	match.SelectChampion("- Xin Yasuo nhaa", jinx)
	match.SelectChampion("- Thua tại Yasuo", thresh)

	match.SelectChampion("- Thảo pick Yasuo", yasuo)
}

