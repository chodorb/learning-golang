package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	stop := false

	for {
		opt, _ := getInput("Choose option (a - add item, s - save bill, t - tip, q - quit)\n", reader)

		switch opt {
		case "a":
			name, _ := getInput("Item name: ", reader)
			priceString, _ := getInput("Item price: ", reader)

			price, err := strconv.ParseFloat(priceString, 64)

			if err != nil {
				fmt.Println("Invalid price")
				break
			}

			fmt.Printf("Name: %v, Price: %v\n", name, price)
			b.addItem(name, price)

		case "s":
			fmt.Println("Saving bill %v", b.format())
			b.save()
			stop = true
		case "t":
			amountString, _ := getInput("Tip amount: ", reader)
			amount, err := strconv.ParseFloat(amountString, 64)
			if err != nil {
				fmt.Println("Invalid amount")
				break
			}

			fmt.Printf("Amount: %v\n", amount)
			b.updateTip(amount)
		case "q":
			stop = true
		default:
			fmt.Println("Invalid Option")
		}
		if stop {
			break
		}
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Bill name: ", reader)
	myBill := newBill(name)

	return myBill
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
