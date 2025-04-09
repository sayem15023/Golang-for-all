package main

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxsum := sum
	for i := k; i < n; i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > maxsum {
			maxsum = sum
		}
	}
	return float64(maxsum) / float64(k)
}

// func main() {
// 	nums1 := []int{1, 12, -5, -6, 50, 3}
// 	k := 4
// 	fmt.Println("maximum average of nums1", findMaxAverage(nums1, k))
// }
