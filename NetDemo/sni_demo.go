package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// 从 TLS 握手消息中提取 SNI 信息
func getSNI(buf []byte) string {
	var sni string
	re := regexp.MustCompile(`^(?:[a-z0-9-]+\.)+[a-z]+$`)

	for i := 0; i < len(buf); i++ {
		if i > 0 && buf[i-1] == 0 && buf[i] == 0 {
			start := i + 2
			length := int(buf[i+1])
			end := start + length
			if start < end && end <= len(buf) {
				str := string(buf[start:end])
				if re.MatchString(str) {
					sni = str
					break
				}
			}
		}
	}

	return sni
}

func main() {
	iface := "lo" // 替换为你的网卡接口名
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置 PacketSource 从网卡读取数据包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// 处理数据包
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer == nil {
			continue
		}

		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer == nil {
			continue
		}

		ip, _ := ipLayer.(*layers.IPv4)
		switch ip.Protocol {
		case layers.IPProtocolTCP:
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)
				fmt.Printf("TCP 数据包: %s:%d -> %s:%d\n", ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)

				if tcp.DstPort == 443 {
					// 提取 TLS 数据包中的 SNI 信息
					sni := getSNI(tcp.Payload)
					if sni != "" {
						fmt.Printf("SNI: %s\n", sni)
					}
				}
			}
		case layers.IPProtocolUDP:
			udpLayer := packet.Layer(layers.LayerTypeUDP)
			if udpLayer != nil {
				udp, _ := udpLayer.(*layers.UDP)
				fmt.Printf("UDP 数据包: %s:%d -> %s:%d\n", ip.SrcIP, udp.SrcPort, ip.DstIP, udp.DstPort)
			}
		}
	}
}
