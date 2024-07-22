package main

import "fmt"

func updateName(x string) {
	// we modify the copy, not the original
	x = "Bazyleusz"
}

func updateClients(c map[int]string) {
	// we modify the variable from the pointers, effecting outer scope
	c[3] = "Andrzej"
}

func main() {
	name := "Bazyli"

	// for non-pointers-wrapper values (string), we pass a copy
	updateName(name)

	// the value does not change
	fmt.Println(name)

	clients := map[int]string{
		1: "Bazyli",
		2: "Kinga",
	}
	// for pointers-wrapper values (maps), we pass the pointers
	updateClients(clients)
	// value changed
	fmt.Println(clients)
}
