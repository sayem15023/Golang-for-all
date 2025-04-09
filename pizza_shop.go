package main

import (
	"fmt"
	"net"
)

func main() {
	// Start a server on port 9000
	listener, _ := net.Listen("tcp", ":9000")
	fmt.Println("Pizza server is running on port 9000...")

	for {
		conn, _ := listener.Accept() // Wait for client connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	order := string(buffer[:n])
	fmt.Println("Customer ordered:", order)

	response := "Your pizza will arrive in 30 minutes!"
	conn.Write([]byte(response))
}
