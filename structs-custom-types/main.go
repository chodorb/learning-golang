package main

import "fmt"

func main() {
	myBill := newBill("Bazyli's bill")

	fmt.Println(myBill.format())
}
