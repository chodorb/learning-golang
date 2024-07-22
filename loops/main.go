package main

import "fmt"

func main() {
	// the while loop
	x := 0
	for x < 5 {
		fmt.Println(x, "is less then 5")
		x++
	}

	// classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i, "is less then 5")
	}

	// for loop for all slice elements (ugly)
	names := []string{"Bazyli", "Kinga", "Marcin", "Kasia"}

	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	// A better for in loop
	for _, value := range names {
		fmt.Println(value)
	}
}
