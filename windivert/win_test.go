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
