package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]

	}
	return strings.Join(words, " ")
}

// func main() {
// 	input1 := "Hello world"
// 	fmt.Println("Actual valu ", input1)
// 	fmt.Println("reverse value", reverseWords(input1))
// }
