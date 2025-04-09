package main

func dominantIndex(nums []int) int {
	max, secondMax, index := 0, 0, 0

	// First, identify the maximum and second maximum elements
	for i, num := range nums {
		if num > max {
			secondMax = max // Update secondMax to the previous max
			max = num       // Update max to the current number
			index = i       // Record the index of the new max
		} else if num > secondMax {
			secondMax = num // Update secondMax if the current number is less than max but greater than secondMax
		}
	}

	// Check if the largest number is at least twice as large as the second largest
	if max >= 2*secondMax {
		return index // Return the index of the largest number if the condition is satisfied
	}
	return -1 // Otherwise, return -1
}

// // main function
// func main() {
// 	// Example input
// 	nums := []int{3, 6, 1, 0}

// 	// Call the dominantIndex function
// 	result := dominantIndex(nums)

// 	// Print the result
// 	fmt.Printf("The dominant index is: %d\n", result)
// }
