package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected deck length of 24, but got %v", len(d))
	} else if d[0] != "Aces of Spades" {
		t.Errorf("Expected first card to be Aces of Spades, but got %v", d[0])
	} else if d[len(d)-1] != "Six of Diamonds" {
		t.Errorf("Expected last card to be Six of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromfile("_decktesting")

	if len(loadedDeck) != 24 {
		t.Errorf("Expected 24 cards in the deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
