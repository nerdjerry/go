package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	prateek := person{firstName: "Prateek", lastName: "Singhal"}
	fmt.Println(prateek)
}
