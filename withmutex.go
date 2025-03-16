package main

import (
	"fmt"
	"sync"
	"time"
)

var availableSeatsInCinema = map[string]bool{"A1": true}
var mu sync.Mutex // Mutex for synchronizing access

func bookSeats(seat string) {
	mu.Lock()         // Lock before accessing shared resource
	defer mu.Unlock() // Unlock when done

	if availableSeatsInCinema[seat] {
		fmt.Println("Booking", seat)
		time.Sleep(1 * time.Second) // Simulate processing delay
		availableSeatsInCinema[seat] = false
		fmt.Println(seat, "booked successfully!")
	} else {
		fmt.Println(seat, "is already booked.")
	}
}

func main() {
	go bookSeats("A1") // User 1 tries to book A1
	go bookSeats("A1") // User 2 also tries to book A1

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final Seat Status:", availableSeatsInCinema)
}
