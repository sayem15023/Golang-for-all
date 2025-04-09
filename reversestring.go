package main

func reverseWord(word string) string {
	reverseword := ""

	for i := len(word) - 1; i >= 0; i-- {
		reverseword += string(word[i])
	}
	return reverseword
}

// func main() {
// 	word := "hello"
// 	reversed := reverseWord(word)
// 	fmt.Println("Actual word ", word)
// 	fmt.Println("Reversed word", reversed)
// }
