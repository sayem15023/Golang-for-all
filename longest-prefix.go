package main

import (
	"sort"
)

func longestCommonPrefic(arr []string) string {
	sort.Strings(arr)
	first := arr[0]
	last := arr[len(arr)-1]
	minlength := min(len(first), len(last))

	i := 0
	for i < minlength && first[i] == last[i] {
		i++
	}
	return first[:i]
}

func mai(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	arr := []string{"geeksforgeeks", "geeks", "geek", "geezer"}
// 	fmt.Println(longestCommonPrefic(arr))
// }
