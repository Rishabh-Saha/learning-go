package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	updateColors(colors)
	printMap(colors)
}
func updateColors(c map[string]string) {
	c["white"] = "#njkcds"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
