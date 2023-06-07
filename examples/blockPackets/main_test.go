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
			log.Print(packet)
			wd.Send(packet)
		}(wd, packet)

	}
}
func TestXxx(t *testing.T) {
	winDivert, err := windivert.NewWinDivertHandle("udp.DstPort==55954")
	if err != nil {
		t.Fatal(err)
	}
	v, err := winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_SIZE)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)

	v, err = winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_LENGTH)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)
	v, err = winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_TIME)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)
	defer winDivert.Close()
	err = winDivert.SetParam(windivert.WINDIVERT_PARAM_QUEUE_SIZE, windivert.WINDIVERT_PARAM_QUEUE_SIZE_MAX)
	if err != nil {
		t.Fatal(err)
	}
	err = winDivert.SetParam(windivert.WINDIVERT_PARAM_QUEUE_LENGTH, windivert.WINDIVERT_PARAM_QUEUE_LENGTH_MAX)
	if err != nil {
		t.Fatal(err)
	}
	err = winDivert.SetParam(windivert.WINDIVERT_PARAM_QUEUE_TIME, windivert.WINDIVERT_PARAM_QUEUE_TIME_MAX)
	if err != nil {
		t.Fatal(err)
	}

	v, err = winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_SIZE)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)

	v, err = winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_LENGTH)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)
	v, err = winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_TIME)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)

	packetChan, err := winDivert.PacketExs()
	if err != nil {
		t.Fatal(err)
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
	v, err := winDivert.GetParam(windivert.WINDIVERT_PARAM_QUEUE_SIZE)
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
	err = winDivert.SetParam(windivert.WINDIVERT_PARAM_QUEUE_SIZE, windivert.WINDIVERT_PARAM_QUEUE_SIZE_MAX)
	if err != nil {
		t.Fatal(err)
	}

}
