package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected Spades of Ace for the first card, but got %v", d[0])
	}
}

// t.Errorf("Expected deck length of 16, but got %v", len(d))
