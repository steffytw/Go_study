package main

import (
	"fmt"
)

func main() {

	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// cards := newCard()
	cards := []string{"Ace of Spades", newCard()}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
