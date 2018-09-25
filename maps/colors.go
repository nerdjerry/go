package main

import "fmt"

func main() {
	var colors map[string]string

	colors = map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	//Literal way to declare map
	//maps := map[string]string{
	//	"test": "value",
	//}

	//Adding new key
	colors["blue"] = "#00ff00"

	printMaps(colors)
}

func printMaps(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex value of", color, "is", hex)
	}
}
