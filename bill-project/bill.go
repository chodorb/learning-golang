package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills function
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// additional () to receive the struct
// it is often good to pass a pointer
// to prevent a lot of copies from being created
func (b *bill) format() string {
	fs := "Bill breakdown:"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("\n%-25v ... $%v", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("\n Tip: %-250.2f", b.tip)
	fs += fmt.Sprintf("\n Total items: %-250.2f", total+b.tip)
	return fs

}

// update the tip
// when we update struct values
// we pass in a pointer
// and go automaticly dereferences it
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill to file
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile(fmt.Sprintf("bills/%v.txt", b.name), data, 0644)
	if err != nil {
		fmt.Println("Error saving bill:", err)
	}
}
