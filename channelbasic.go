package main

import "fmt"

// sender and receiver in an unbuffered channel
func main() {
	ch := make(chan string) //create an unbuffer channel
	go func() {
		ch <- "hello from sender"
	}()
	msg := <-ch

	fmt.Println(msg)
}
