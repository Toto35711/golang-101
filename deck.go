package main

import "fmt"

// go's style of "OOP", it extends []string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// you can think of it like a method of the deck class
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
