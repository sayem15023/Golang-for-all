package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {

		if flowerbed[i] == 0 {
			leftempty := (i == 0 || flowerbed[i-1] == 0)
			rightempty := (i == len(flowerbed)-1 || flowerbed[i+1] == 0)

			if leftempty && rightempty {
				flowerbed[i] = 1
				n--

				if n == 0 {
					return true
				}
			}
		}
	}
	return n <= 0
}

// func main() {
// 	// Test cases
// 	flowerbed1 := []int{0, 1, 0, 1, 1}
// 	n1 := 1
// 	fmt.Println(canPlaceFlowers(flowerbed1, n1)) // Output: true

// 	flowerbed2 := []int{1, 0, 1}
// 	n2 := 2
// 	fmt.Println(canPlaceFlowers(flowerbed2, n2)) // Output: false
// }
