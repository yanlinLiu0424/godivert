package windivert

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"testing"
	"unsafe"
)

type tt struct {
	name string
}

func TestXxx(t *testing.T) {
	cs := make(chan []tt, 1)

	//cs = append(cs, tt{name: "test123"})
	for c := range cs {
		log.Print(c)
	}
}

func TestX(t *testing.T) {

	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(float64(0)))

}

func TestXXX(t *testing.T) {
	packets := []*Packet{}
	for i := 0; i <= 100; i++ {
		p := &Packet{}
		packets = append(packets, p)
	}
	log.Print(len(packets))
}

func TestXxx1(t *testing.T) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		fmt.Printf("interface:%v,name:%v\n", i.Index, i.Name)
	}

}
func TestXx1x(t *testing.T) {
	//addr := make([]Address, quantity)
	b := []byte{0x00, 0x70}
	v := int(binary.BigEndian.Uint16(b))
	log.Print(v)
}
