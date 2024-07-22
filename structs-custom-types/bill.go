package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills function
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"pie":  21.99,
			"cola": 2.99,
		},
		tip: 0,
	}
	return b
}

// additional () to receive the struct
func (b bill) format() string {
	fs := "Bill breakdown:"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("\n%-25v ... $%v", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("\n Total items: %-250.2f", total)
	return fs
}
