package main

import (
	"os"
	"testing"
)

var numberOfCards = 16

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckSize := len(d)

	if deckSize != numberOfCards {
		t.Errorf("Expected deck length of %v, but got %v", numberOfCards, deckSize)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[deckSize-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[deckSize-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting.txt"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)
	loadedDeckSize := len(loadedDeck)

	if loadedDeckSize != numberOfCards {
		t.Errorf("Expected %v cards in deck, got %v", numberOfCards, loadedDeckSize)
	}

	os.Remove(fileName)
}
