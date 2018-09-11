package main

func main() {
	//Create a slice of string
	cards := deck{"Ace of Diamonds", getCard()}
	//Add value to slice and expand its capacity
	cards = append(cards, "6 of Club")

	cards.print()
}

func getCard() string {
	return "Ace of Spades"
}
