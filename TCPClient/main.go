// tcp_client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080") // Connect to server
	defer conn.Close()

	conn.Write([]byte("Hello from client!\n")) // Send message

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf) // Read response
	fmt.Println("Received from server:", string(buf[:n]))
}
