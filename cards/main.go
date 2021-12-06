package main

import "fmt"

func main() {
	fmt.Println("init")
	cards := newDeck()
	cards.print()

	fmt.Println("shuffle")
	cards.shuffle()
	cards.print()

	hand, afterDeck := deal(cards, 5)
	fmt.Println("deal hand")
	hand.print()

	fmt.Println("deal deck")
	afterDeck.print()

	fmt.Println("read from file")
	cards.saveToFile("cards")
	c := readFromFile("cards")
	c.print()
}
