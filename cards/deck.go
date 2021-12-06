package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) shuffle() {
	rand := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func deal(d deck, length int) (deck, deck) {
	return d[:length], d[length:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) {
	s := strings.Join(d, ",")

	ioutil.WriteFile(filename, []byte(s), 0666)
}

func readFromFile(filename string) deck {
	bs, _ := ioutil.ReadFile(filename)

	return strings.Split(string(bs), ",")
}
