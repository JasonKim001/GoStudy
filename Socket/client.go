package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting: ", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connected to server.")

	reader := bufio.NewReader(os.Stdin)
	for {
		// read message from stdin
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// send message to server
		fmt.Fprintf(conn, message)

		// read reply from server
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Response from server: ", response)
	}
}
