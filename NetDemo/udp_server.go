package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	log.Println("UDP server listening on 127.0.0.1:8080")

	for {
		handleUDPConnection(conn)
	}
}

func handleUDPConnection(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Printf("Failed to read data: %v", err)
		return
	}

	log.Printf("Received from %s: %s", addr, string(buf[:n]))

	// Echo back the message
	message := fmt.Sprintf("Hello, %s", addr.String())
	_, err = conn.WriteToUDP([]byte(message), addr)
	if err != nil {
		log.Printf("Failed to send response: %v", err)
		return
	}
}
