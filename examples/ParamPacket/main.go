package main

import (
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/yanlinLiu0424/godivert/windivert"
)

func checkPacketEx(wd *windivert.WinDivertHandle, packetChan <-chan *windivert.Packet) {
	for packet := range packetChan {
		go func(wd *windivert.WinDivertHandle, packet *windivert.Packet) {
			//wd.HelperParsePacket(packet.Raw)
			/*srcip := packet.SrcIP().String()
			destip := packet.DstIP().String()
			srp, _ := packet.SrcPort()
			dp, _ := packet.DstPort()
			fmt.Print(srcip, destip, srp, dp)*/
			_, err := wd.HelperParsePacket(packet.Raw)
			if err == nil {
				log.Printf("srcip:%v dstip:%v ", packet.SrcIP(), packet.DstIP())

			}
			wd.Send(packet)
		}(wd, packet)

	}
}
func main() {
	winDivert, err := windivert.NewWinDivertHandle("true")
	if err != nil {
		panic(err)
	}

	defer winDivert.Close()
	/*	err = winDivert.SetParam(windivert.WinDivertParamQueueSize, windivert.WinDivertParamQueueSizeMax)
		if err != nil {
			panic(err)
		}
		err = winDivert.SetParam(windivert.WinDivertParamQueueLength, windivert.WinDivertParamQueueLengthMax)
		if err != nil {
			panic(err)
		}
		err = winDivert.SetParam(windivert.WinDivertParamQueueTime, windivert.WinDivertParamQueueTimeMax)
		if err != nil {
			panic(err)
		}*/

	packetChan, err := winDivert.PacketExs()
	if err != nil {
		panic(err)
	}

	go checkPacketEx(winDivert, packetChan)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func TestGetParam(t *testing.T) {
	winDivert, err := windivert.NewWinDivertHandle("icmp")
	if err != nil {
		t.Fatal(err)
	}
	defer winDivert.Close()
	v, err := winDivert.GetParam(windivert.WinDivertParamQueueSize)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)
}
func TestSetParam(t *testing.T) {
	winDivert, err := windivert.NewWinDivertHandle("true")
	if err != nil {
		t.Fatal(err)
	}
	defer winDivert.Close()
	err = winDivert.SetParam(windivert.WinDivertParamQueueSize, windivert.WinDivertParamQueueSizeMax)
	if err != nil {
		t.Fatal(err)
	}

}
