package main

import (
	"fmt"
	"strings"
)

// simple function example
func sayHelloWorld(name string) {
	fmt.Printf("Hello world from %v \n", name)
}

// function to capitalize all names int the slice
func capitalizeSlice(slice []string) []string {
	for index, _ := range slice {
		slice[index] = strings.ToUpper(slice[index])
	}
	return slice
}

// function taking a slice and function as args
func cycleNames(namesSlice []string, f func(string)) {
	namesSlice = capitalizeSlice(namesSlice)
	for _, value := range namesSlice {
		f(value)
	}
}

func main() {
	names := []string{"Bazyli", "Kinga", "Andrzej"}
	cycleNames(names, sayHelloWorld)
}
