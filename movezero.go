package main

func moveZeroes(num []int) {
	left := 0
	for right := 0; right < len(num); right++ {
		if num[right] != 0 {
			if left != right {
				num[left], num[right] = num[right], num[left]
			}
			left++
		}

	}
}

// func main() {
// 	nums1 := []int{10, 15, 0, 22, 34}
// 	moveZeroes(nums1)
// 	fmt.Println(nums1) // Output: [1, 3, 12, 0, 0]

// 	nums2 := []int{0}
// 	moveZeroes(nums2)
// 	fmt.Println(nums2) // Output: [0]
// }
