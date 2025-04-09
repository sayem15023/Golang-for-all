package main

func kidsWithCandies(candies []int, extracandies int) []bool {
	//find the maximum number of candies among all kids

	maxcandies := candies[0]
	for _, candy := range candies {
		if candy > maxcandies {
			maxcandies = candy
		}
	}

	//create a result to store boolean result

	result := make([]bool, len(candies))
	//check if each kid can have the greatest number of candies
	for i, candy := range candies {
		if candy+extracandies >= maxcandies {
			result[i] = true
		} else {
			result[i] = false
		}

	}
	//return the result of array
	return result

}

// func main() {
// 	// Example 1
// 	candies1 := []int{2, 3, 5, 1, 3}
// 	extraCandies1 := 3
// 	fmt.Println(kidsWithCandies(candies1, extraCandies1)) // Output: [true true true false true]

// 	// Example 2
// 	candies2 := []int{4, 2, 1, 1, 2}
// 	extraCandies2 := 1
// 	fmt.Println(kidsWithCandies(candies2, extraCandies2)) // Output: [true false false false false]

// 	// Example 3
// 	candies3 := []int{2, 3, 4, 5}
// 	extraCandies3 := 5
// 	fmt.Println(kidsWithCandies(candies3, extraCandies3)) // Output: [true false true]
// }
