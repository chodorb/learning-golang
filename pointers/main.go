package main

import "fmt"

func updateName(x string) {
	x = "new value"
}

func actuallyUpdateName(x *string) {
	// we set an actual memory value
	// we get the memory location from the pointer
	// &var -> pointer
	// *pointer -> var memory location
	*x = "new value"
}

func main() {
	name := "some value"
	// value does not change
	updateName(name)

	//pointer to memory location
	m := &name
	// dereference to memory location
	fmt.Println(*m)

	actuallyUpdateName(&name)
	// now the value changes
	fmt.Println(name)

}
