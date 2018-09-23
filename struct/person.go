package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var prateek person
	prateek = person{
		firstName: "Prateek",
		lastName:  "Singhal",
		contact: contactInfo{
			email:   "psinghal@gmail.com",
			zipcode: 12344,
		},
	}
	prateek.updateName("Nerd")
	prateek.print()
}

func (p *person) updateName(updatedName string) {
	(*p).firstName = updatedName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
