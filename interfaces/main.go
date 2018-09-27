package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	greeting(eb)
	greeting(sb)
}

func greeting(b bot) {
	fmt.Println(b.getGreeting())
}

//Same logic as english bot, just argument type different
//func getGreeting(sb spanishBot) {
//	fmt.Println(sb.greeting())
//}
func (englishBot) getGreeting() string {
	//Custom logic only for english
	return "Hi, There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
