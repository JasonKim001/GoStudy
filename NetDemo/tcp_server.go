package main

import (
	"crypto/tls"
	"log"
	"net"
)

func main() {
	// 配置 TLS
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load key pair: %v", err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	// 创建监听器
	listener, err := tls.Listen("tcp", "127.0.0.1:443", config)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Println("TLS server listening on 127.0.0.1:443")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleTLSConnection(conn)
	}
}

func handleTLSConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Failed to read data: %v", err)
		return
	}

	log.Printf("Received: %s", string(buf[:n]))
}
