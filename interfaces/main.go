package main

import (
	"fmt"
)

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	//sb := spanishBot{}
	getGreeting(eb)
}

func getGreeting(eb englishBot) {
	fmt.Println(eb.greeting())
}

//Same logic as english bot, just argument type different
//func getGreeting(sb spanishBot) {
//	fmt.Println(sb.greeting())
//}
func (englishBot) greeting() string {
	//Custom logic only for english
	return "Hi, There!"
}

func (spanishBot) greeting() string {
	return "Hola!"
}
