package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be 'Ace of Spades' but got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card in deck to be 'King of Clubs' but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	deckToSave := newDeck()
	deckToSave.saveToFile(filename)

	deckRead := newDeckFromFile(filename)
	if len(deckRead) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(deckRead))
	}

	os.Remove(filename)
}
