package main

func main() {
	cards := newDeck()

	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()
	//cards.writeToFile("deck")

	cards = readFromFile("deck")
	cards.print()
}
