package main

//Working with array and slice

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// hand, remainingCard := deal(cards, 7)

	// hand.print()
	// remainingCard.print()
	// cards.saveToFile("my_cards")
}
func newCard() string {
	return "Not ace card"
}
