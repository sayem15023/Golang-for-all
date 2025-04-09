package main

func removeStars(s string) string {
	stack := []rune{}

	for _, char := range s {
		if char == '*' {

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {

			stack = append(stack, char)
		}
	}
	return string(stack)
}

// func main() {
// 	// Example 1
// 	s1 := "leet**cod*e"
// 	fmt.Println(removeStars(s1)) // Output: "lecoe"

// 	// Example 2
// 	s2 := "erase*****"
// 	fmt.Println(removeStars(s2)) // Output: ""
// }
