package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听指定的端口
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listeninig: ", err)
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// 接受客户端连接
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		fmt.Println("Client connected.")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		// read data from client
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading: ", err)
				return
			}
			fmt.Print("Message received: ", message)

			// send reply to client
			conn.Write([]byte("Message received.\n"))
		}
	}
}
