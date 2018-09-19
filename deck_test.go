package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 24 {
		t.Errorf("Expected 30 cards but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Five of Clubs" {
		t.Errorf("Expected Five of Clubs but got %v", deck[len(deck)-1])
	}
}
