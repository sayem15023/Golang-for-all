package main

import (
	"fmt"
)

func main() {
	price := 100.0
	discount := 20.0 // percent

	finalPrice := applyDiscount(price, discount)

	fmt.Printf("Final price after %.2f%% discount is: %.2f\n", discount, finalPrice)
}

func applyDiscount(price float64, discountPercent float64) float64 {
	discountAmount := price * discountPercent / 100
	return price - discountAmount
}
