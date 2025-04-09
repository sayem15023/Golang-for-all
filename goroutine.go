package main

import (
	"fmt"
	"time"
)

func takeOrder(orderId int) {
	fmt.Printf("Processing order #%d\n", orderId)
	time.Sleep(2 * time.Second) // Simulate cooking time
	fmt.Printf("Order #%d is ready!\n", orderId)
}

func main() {
	for i := 1; i < 30; i++ {
		go takeOrder(i)
	}
	time.Sleep(5 * time.Second) // Give enough time for goroutines to complete
	fmt.Println("All orders processed!")
}
