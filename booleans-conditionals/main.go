package main

import "fmt"

func main() {
	age := 45

	// Some boolean expression
	fmt.Println(age > 50)

	// example if + elif + else
	if age < 30 {
		fmt.Printf("Age %v is less than 30\n", age)
	} else if age < 40 {
		fmt.Printf("Age %v is greater than 40\n", age)
	} else {
		fmt.Printf("Age %v is big\n", age)
	}

	// continue and breaking
	names := []string{"Alice", "Bob", "Mario", "Luigi", "Yoshi"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continue at pos", index)
			// continue jumps to loop block beginning
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("%v is %v\n", index, value)
	}
}
