package main

func isSubsequence(s string, t string) bool {

	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)

}

// func main() {
// 	s1 := "abcsssss"
// 	t1 := "axbd"
// 	fmt.Println(isSubsequence(s1, t1))
// }
