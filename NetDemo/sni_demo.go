package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var udpCount = 0
var tcpCount = 0
var sniCount = 0

// Function to extract SNI from TLS handshake
func getSNI(data []byte) string {
	if len(data) < 5 || data[0] != 0x16 {
		return ""
	}
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			return ""
		}
		if data[i] == 0x00 && data[i+1] == 0x17 {
			j := i + 5
			for j < len(data)-1 {
				if data[j] == 0x00 {
					break
				}
				j++
			}
			return string(data[i+5 : j])
		}
	}
	return ""
}

func main() {
	iface := "eth0" // Replace with your network interface
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set up a packet source to read packets from the handle
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// Process packet here
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
				if tcpCount == 0 {
					fmt.Printf("TCP Packet: %s:%d -> %s:%d\n", ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)
					tcpCount += 1
				}

				if tcp.DstPort == 443 {
					sni := getSNI(tcp.Payload)
					if sni != "" {
						if sniCount == 0 {
							fmt.Printf("SNI: %s\n", sni)
							sniCount += 1
						}

					}
				}
			}
		case layers.IPProtocolUDP:
			udpLayer := packet.Layer(layers.LayerTypeUDP)
			if udpLayer != nil {
				if udpCount == 0 {
					udp, _ := udpLayer.(*layers.UDP)
					fmt.Printf("UDP Packet: %s:%d -> %s:%d\n", ip.SrcIP, udp.SrcPort, ip.DstIP, udp.DstPort)
					udpCount += 1
				}
			}
		}
	}
}
