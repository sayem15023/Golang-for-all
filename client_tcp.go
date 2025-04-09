package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:9000") // Connect to server

	order := "2 Large Pepperoni Pizza"
	conn.Write([]byte(order)) // Send the order

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer) // Get response from server
	fmt.Println("Pizza Shop:", string(buffer[:n]))
}
