package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("NewDeck returned wrong length: got %d, want 16", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("NewDeck returned wrong value: got %v, want Ace of Spades", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("NewDeck returned wrong value: got %v, want Four of Clubs", d[len(d)-1])
	}
}
