package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func gcdOfString(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len1 := len(str1)
	len2 := len(str2)
	gcdlen := gcd(len1, len2)
	return str1[:gcdlen]
}

// func main() {
// 	// Example 1
// 	str1 := "ABCABC"
// 	str2 := "ABC"
// 	fmt.Println("Example 1:")
// 	fmt.Printf("Input: str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
// 	fmt.Printf("Output: \"%s\"\n", gcdOfString(str1, str2))
// 	fmt.Println()

// 	// Example 2
// 	str1 = "ABABAB"
// 	str2 = "ABAB"
// 	fmt.Println("Example 2:")
// 	fmt.Printf("Input: str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
// 	fmt.Printf("Output: \"%s\"\n", gcdOfString(str1, str2))
// 	fmt.Println()

// 	// Example 3
// 	str1 = "LEET"
// 	str2 = "CODE"
// 	fmt.Println("Example 3:")
// 	fmt.Printf("Input: str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
// 	fmt.Printf("Output: \"%s\"\n", gcdOfString(str1, str2))
// 	fmt.Println()

// 	// Additional Test Cases
// 	fmt.Println("Additional Test Cases:")

// 	// Case where both strings are the same
// 	str1 = "AAAAAA"
// 	str2 = "AAAAAA"
// 	fmt.Printf("Input: str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
// 	fmt.Printf("Output: \"%s\"\n", gcdOfString(str1, str2))
// 	fmt.Println()

// 	// Case where one string is a single character
// 	str1 = "A"
// 	str2 = "AAAAA"
// 	fmt.Printf("Input: str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
// 	fmt.Printf("Output: \"%s\"\n", gcdOfString(str1, str2))
// 	fmt.Println()

// 	// Case where no common divisor exists
// 	str1 = "XYZ"
// 	str2 = "ABC"
// 	fmt.Printf("Input: str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
// 	fmt.Printf("Output: \"%s\"\n", gcdOfString(str1, str2))
// 	fmt.Println()
// }
