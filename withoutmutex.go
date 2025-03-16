package main

import (
	"fmt"
	"time"
)

var availableSeats = map[string]bool{"A1": true} // Shared memory (seat status)

func bookSeat(seat string) {
	if availableSeats[seat] { // Check if seat is available
		fmt.Println("Booking", seat)
		time.Sleep(1 * time.Second)  // Simulating delay in processing
		availableSeats[seat] = false // Mark seat as booked
		fmt.Println(seat, "booked successfully!")
	} else {
		fmt.Println(seat, "is already booked.")
	}
}

func main() {
	go bookSeat("A1") // User 1 tries to book A1
	go bookSeat("A1") // User 2 also tries to book A1

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final Seat Status:", availableSeats)
}
