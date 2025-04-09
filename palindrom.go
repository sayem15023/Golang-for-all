package main

func palinDrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true

}

// func main() {
// 	word := "radar"
// 	if palinDrome(word) {
// 		fmt.Println(word, "word is palindrom")
// 	} else {
// 		fmt.Println(word, "word is not palindrom")
// 	}
// }
