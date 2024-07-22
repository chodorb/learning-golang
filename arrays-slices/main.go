package main

import "fmt"

func main() {
	// define an array with 5 int elements
	// var ages [5]int = [5]int{18,19,20,21,22}
	// var ages = [5]int{18,19,20,21,22}
	ages := [5]int{18, 19, 20, 21, 22}
	fmt.Println(ages, len(ages))
	
	// slices dont need fixed lengths (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 80
	// slice.append() kind of stuff
	scores = append(scores, 70)

	fmt.Println(scores, len(scores))

	// slice ranges <1:3) range
	rangeOne := scores[1:3]
	fmt.Println(rangeOne)

	// slice ranges <2:end>
	rangeTwo := scores[2:]
	fmt.Println(rangeTwo)

	// slice ranges <start:3)
	rangeThree := scores[:3]
	fmt.Println(rangeThree)
}
