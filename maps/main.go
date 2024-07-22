package main

import "fmt"

func main() {
	// a simple map, with string keys and float values
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
	}

	// add new key value pair
	menu["burger"] = 10.99

	// loop through a map
	for key, value := range menu {
		fmt.Println(key, value)
	}

	// delete some element
	delete(menu, "burger")

}
