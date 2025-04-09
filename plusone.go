package main

func plusOne(digits []int) []int {
	// Start from the last digit and move towards the first
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			// If the current digit is less than 9, increment it and return
			digits[i]++
			return digits
		}
		// If the digit is 9, set it to 0 and continue
		digits[i] = 0
	}

	// If all digits were 9, we need an additional digit at the front
	// Example: 999 + 1 -> 1000
	return append([]int{1}, digits...)
}

// func main() {
// 	digits := []int{9, 9, 9}
// 	result := plusOne(digits)
// 	fmt.Println(result) // Output: [1, 2, 4]
// }
