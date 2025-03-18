package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10 // Sender sends value 10
	ch <- 20 // Sender sends value 20 (buffer is now full)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
