package main

import (
	"fmt"
	"strings"
)

// function to return 2 return values
func getInitials(str string) (string, string) {
	str = strings.ToUpper(str)
	strSlice := strings.Split(str, " ")

	var initials []string

	for _, v := range strSlice {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	s1 := "Bazyli Chodor"
	initial1, initial2 := getInitials(s1)
	fmt.Println(initial1, initial2)
}
