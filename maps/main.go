package main

import "fmt"

func main() {

	//How to declare maps
	//A map of string, string
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}

	// Declaring map without key:value pair
	//go will init map with it's empty/zero value
	var colorMap map[string]string

	//Yet another way to declare map
	hex := make(map[string]string)

	//how to populate a key/value in map
	hex["#0000FF"] = "blue"

	fmt.Println(colors)
	fmt.Println(colorMap)
	fmt.Println(hex)
	printMap(colors)
}

//Print the given map
func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("value for key [", color, "] is [", hex, "]")
	}
}
