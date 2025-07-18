package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 4 {
		t.Errorf("Expected deck length of 4, but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%s'", d[0])
	}

}
