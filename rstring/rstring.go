package main

import (
	"fmt"
)

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var inputString string
	fmt.Scanln(&inputString)
	fmt.Printf("%s\n", Reverse(inputString))
}
