package main

import (
	"log"
	"net"
	"time"

	"github.com/yanlinLiu0424/godivert/windivert"
)

var cloudflareDNS = net.ParseIP("8.8.4.4")

func checkPacket(wd *windivert.WinDivertHandle, packetChan <-chan *windivert.Packet) {
	for packet := range packetChan {
		if !packet.DstIP().Equal(cloudflareDNS) {
			log.Print(packet)
			packet.Send(wd)
		}
	}
}

func main() {
	winDivert, err := windivert.NewWinDivertHandle("icmp")
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packetChan, err := winDivert.Packets()
	if err != nil {
		panic(err)
	}

	go checkPacket(winDivert, packetChan)

	time.Sleep(1 * time.Minute)
}
