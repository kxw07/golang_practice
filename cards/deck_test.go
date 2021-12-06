package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Deck init length is not 16, got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Deck init first card is not Ace of Spades, got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Deck init last card is not Four of Clubs, got %v", deck[len(deck)-1])
	}
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	numberOfDealCards := 5
	handCards, leftDeck := deal(deck, numberOfDealCards)

	if len(handCards) != numberOfDealCards {
		t.Errorf("It should be %v cards in hand, but got %v", numberOfDealCards, len(handCards))
	}

	if len(leftDeck) != (len(deck) - numberOfDealCards) {
		t.Errorf("It should left %v in deck, but got %v", (len(deck) - numberOfDealCards), len(leftDeck))
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_cardsTesting")
	deck := newDeck()
	deck.saveToFile("_cardsTesting")

	newDeck := readFromFile("_cardsTesting")
	if len(newDeck) != 16 {
		t.Errorf("It should be 16 cards in file, but got %v", len(newDeck))
	}

	os.Remove("_cardsTesting")
}
