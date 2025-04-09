package main

func uniqueOccurrences(arr []int) bool {
	freqMap := make(map[int]int)

	for _, nums := range arr {
		freqMap[nums]++
	}

	seen := make(map[int]bool)

	for _, fre := range freqMap {
		if seen[fre] {
			return false
		}
		seen[fre] = true
	}
	return true
}

// func main() {
// 	arr := []int{1, 2, 2, 1, 3, 3}
// 	fmt.Println(uniqueOccurrences(arr))
// }
