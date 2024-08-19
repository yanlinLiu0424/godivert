package windivert

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestIPv6(t *testing.T) {
	hexStr := "6e00000000200001fe8000000000000002d6fefffee7d019ff0200000000000000000001ffe7d0193a0005020000010083000f4900000000ff0200000000000000000001ffe7d019"
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatal(err)
	}
	addr := Packet{Raw: data}
	log.Print(addr.SrcIP())
	log.Print(addr.DstIP())
	log.Print(addr.String())
}
func TestIPv61(t *testing.T) {
	data := []byte{
		0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
	}
	addr := Packet{Raw: data}
	log.Print(addr.SrcIP())
	log.Print(addr.DstIP())
	log.Print(addr.String())
}

func TestT1n(t *testing.T) {
	v := "6000000000240001fe80000000000000bd3b48b576234f95ff0200000000000000000000000000163a00050200000100"
	data, err := hex.DecodeString(v)
	if err != nil {
		t.Fatal(err)
	}
	addr := Packet{Raw: data}
	log.Print(addr.SrcIP())
	log.Print(addr.DstIP())
	log.Print(addr.String())
}
