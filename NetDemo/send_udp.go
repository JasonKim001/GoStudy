package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// 目标地址
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to resolve server address: %v", err)
	}

	// 创建 UDP 连接
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 发送 UDP 报文
	message := "Hello, UDP server"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("UDP message sent")
}
