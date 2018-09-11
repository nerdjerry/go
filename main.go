package main

import "fmt"

func main() {
	//Create a slice of string
	cards := []string{"Ace of Diamonds", getCard()}
	//Add value to slice and expand its capacity
	cards = append(cards, "6 of Club")

	fmt.Println(cards)
}

func getCard() string {
	return "Ace of Spades"
}
