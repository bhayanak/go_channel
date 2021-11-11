package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	totalCards := 52
	if len(d) != totalCards {
		t.Errorf("Expected %d but got %d", totalCards, len(d))
	}
}
