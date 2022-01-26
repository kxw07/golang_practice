package main

import "fmt"

type Bot interface {
	getWords() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	greeting(eb)
	greeting(sb)
}

func greeting(b Bot) {
	fmt.Println(b.getWords())
}

func (eb EnglishBot) getWords() string {
	return "Hello!"
}

func (sb SpanishBot) getWords() string {
	return "Hola!"
}
