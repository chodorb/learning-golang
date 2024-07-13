package main

import "fmt"

func main() {
	// set type manually
	var stringVariable string = "XD"
	fmt.Println(stringVariable)

	// automaticly infer the type
	var stringVariable2 = "XDD"
	fmt.Println(stringVariable2)

	// declare variable without value
	var stringVariable3 string
	fmt.Println(stringVariable3)

	// declare a variable with a type (only in functioncs)
	stringVariable4 := "XDDD"
	fmt.Println(stringVariable4)

	// integers
	var intVariable int = 20
	var intVariable2 = 30
	intVariable3 := 40

	fmt.Println(intVariable + intVariable2 + intVariable3)

	// bits and memory
	var num1 int8 = 127
	fmt.Println(num1)

	// ints > 0
	var num2 uint = 20
	fmt.Println(num2)

	// floats must be assigned bit size
	var float1 float32 = 2.5
	var float2 float64 = 321312321.213

	fmt.Println(float1)
	fmt.Println(float2)

}
