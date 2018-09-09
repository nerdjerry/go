package main

import "fmt"

func main() {
	//var card string = "Aces of Spades"
	card := getCard()

	fmt.Println(card)
}

func getCard() string {
	return "Ace of Spades"
}
