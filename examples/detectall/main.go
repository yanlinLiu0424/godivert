package main

import (
	"os"
	"os/signal"

	"github.com/yanlinLiu0424/godivert/windivert"
)

func checkPacket(wd *windivert.WinDivertHandle, packetChan <-chan *windivert.Packet) {
	for packet := range packetChan {
		go func(wd *windivert.WinDivertHandle, packet *windivert.Packet) {
			//raw, _ := wd.HelperParsePacket(packet.Raw)
			//log.Print(raw)
			packet.Send(wd)
		}(wd, packet)

	}
}

func main() {
	winDivert, err := windivert.NewWinDivertHandle("!loopback")
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packetChan, err := winDivert.PacketExs()
	if err != nil {
		panic(err)
	}

	go checkPacket(winDivert, packetChan)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
