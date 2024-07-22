package main

import "fmt"

func main() {
	// Print line function, prints stuff in newline
	fmt.Println("Hello World!")

	// Print function, prints with not newline
	fmt.Print("Hello ")
	fmt.Print("World\n")


	// The fstring equivalent
	var age int = 23
	var name string = "Bazyli"
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	// %_ = format specifier
	// %v is a variable
	// %q is quoted variable
	// %T is a type
	// %0.2f is a float ronded to 2 decimal points


	// Here we can save the fstring
	var str string = fmt.Sprintf("My age is %v", age)
	fmt.Println(str)
}