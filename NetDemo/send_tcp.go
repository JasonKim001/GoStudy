package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// 读取自签名证书
	certPool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Failed to read server certificate: %v", err)
	}
	certPool.AppendCertsFromPEM(serverCert)

	// 配置 TLS 连接，设置 SNI 和自签名证书
	config := &tls.Config{
		ServerName: "example.com", // 设置 SNI
		RootCAs:    certPool,      // 添加自签名证书到证书池
	}

	// 创建一个 TCP 连接到回环地址
	conn, err := net.Dial("tcp", "127.0.0.1:443")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 将 TCP 连接升级为 TLS 连接
	tlsConn := tls.Client(conn, config)

	// 执行握手
	err = tlsConn.Handshake()
	if err != nil {
		log.Fatalf("TLS handshake failed: %v", err)
	}

	// 握手成功后，获取连接状态
	state := tlsConn.ConnectionState()
	fmt.Printf("TLS Handshake succeeded\n")
	fmt.Printf("Server Name: %s\n", state.ServerName)
	fmt.Printf("Cipher Suite: %s\n", tls.CipherSuiteName(state.CipherSuite))
	fmt.Printf("Negotiated Protocol: %s\n", state.NegotiatedProtocol)

	// 发送数据
	message := "Hello, TLS server"
	_, err = tlsConn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}
