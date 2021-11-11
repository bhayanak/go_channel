package main

import "fmt"

func main() {
	// filename := "cards.csv"
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	fmt.Println( " And the rest is :")
	// remainingDeck.print()
}
