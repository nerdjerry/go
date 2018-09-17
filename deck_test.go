package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 24 {
		t.Errorf("Expected 30 cards but got %v", len(deck))
	}
}
