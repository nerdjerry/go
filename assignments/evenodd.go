package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Printf("%v is even", value)
			fmt.Println()
		} else {
			fmt.Printf("%v is odd", value)
			fmt.Println()
		}
	}
}
