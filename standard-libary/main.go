package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends"
	// Some example strings.methods
	fmt.Println(strings.Contains(greeting, "hello"))
	// replace leaves string unchanged
	fmt.Println(strings.ReplaceAll(greeting, "hello", "HELLO"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "hello"))
	fmt.Println(strings.Split(greeting, " "))

	// Some example sort methods
	// sorting alters the original value
	ages := []int{30, 40, 60, 70, 50}
	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(sort.SearchInts(ages, 50))

	names := []string{"Bazyli", "Kinga", "Marcin", "Kasia"}
	sort.Strings(names)
	fmt.Println(sort.SearchStrings(names, "Kasia"))
}
