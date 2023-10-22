package main

func main() {
	cards := newDeck()
	// cards.saveToFile("my_cards")
	// d := newDeckFromFile("my_cards")
	// d.print()
	cards.print()
	cards.shuffle()
	cards.print()
}
