package main

func main() {
	cards := newDeck()
	hands, remainingHands := deal(cards, 5)
	hands.print()
	remainingHands.print()
	cards.print()
}
