package main

import "fmt"

func main() {

	prices := []float64{100.0, 200.0, 50.0}
	discount := 10.0
	tax := 5.0

	finalAmount := calculationFinalAmount(prices, discount, tax)

	fmt.Println("the result is ", finalAmount)

}

func calculationFinalAmount(prices []float64, discount float64, tax float64) float64 {
	total := 0.0

	for _, price := range prices {
		total += price
	}
	discountPrice := (total * discount) / 100
	taxprice := ((total - discountPrice) * tax) / 100
	finalresult := (total - discountPrice) + taxprice

	return finalresult

}
