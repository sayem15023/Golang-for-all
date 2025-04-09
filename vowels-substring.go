package main

func maxVowels(s string, k int) int {
	isVowels := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}
	max_count := 0
	current_count := 0

	for i := 0; i < k; i++ {
		if isVowels(s[i]) {
			current_count++
		}
	}
	max_count = current_count
	for i := k; i < len(s); i++ {
		if isVowels(s[i]) {
			current_count++
		}
		if isVowels(s[i-k]) {
			current_count--
		}
		if current_count > max_count {
			max_count = current_count

		}
	}
	return max_count

}

// func main() {
// 	fmt.Println(maxVowels("sayrty", 3))
// }
