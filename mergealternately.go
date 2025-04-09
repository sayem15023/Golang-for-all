package main

func mergeAlternately(word1 string, word2 string) string {
	var merged []byte // Step 1: Initialize result storage
	i, j := 0, 0      // Step 2: Start indices for both strings

	// Step 3: Traverse both strings alternately
	for i < len(word1) && j < len(word2) {
		merged = append(merged, word1[i]) // Add character from word1
		merged = append(merged, word2[j]) // Add character from word2
		i++
		j++
	}

	// Step 4: Append remaining characters from either string
	if i < len(word1) {
		merged = append(merged, word1[i:]...)
	}
	if j < len(word2) {
		merged = append(merged, word2[j:]...)
	}

	// Step 5: Convert to string and return
	return string(merged)
}

// func main() {
// 	// Example 1
// 	word1 := "abc"
// 	word2 := "pqr"
// 	fmt.Println(mergeAlternately(word1, word2)) // Output: "apbqcr"

// 	// Example 2
// 	word1 = "ab"
// 	word2 = "pqrs"
// 	fmt.Println(mergeAlternately(word1, word2)) // Output: "apbqrs"

// 	// Example 3
// 	word1 = "abcd"
// 	word2 = "pqqwrtiooo"
// 	fmt.Println(mergeAlternately(word1, word2)) // Output: "apbqcd"
// }
