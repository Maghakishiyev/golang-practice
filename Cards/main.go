package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("myCards")
	// fmt.Println([]byte(hand))
	// cards := newDeckFromFile("myCards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
