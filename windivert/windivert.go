package windivert

import (
	"encoding/binary"
	"errors"
	"runtime"
	"syscall"
	"unsafe"
)

var (
	winDivertDLL *syscall.LazyDLL

	winDivertOpen                *syscall.LazyProc
	winDivertClose               *syscall.LazyProc
	winDivertRecv                *syscall.LazyProc
	winDivertRecvEx              *syscall.LazyProc
	winDivertSend                *syscall.LazyProc
	winDivertSendEx              *syscall.LazyProc
	winDivertHelperCalcChecksums *syscall.LazyProc
	winDivertHelperEvalFilter    *syscall.LazyProc
	winDivertHelperCheckFilter   *syscall.LazyProc
	winDivertGetParam            *syscall.LazyProc
	winDivertSetParam            *syscall.LazyProc
	winDivertHelperParsePacket   *syscall.LazyProc
)

func init() {
	LoadDLL("WinDivert.dll", "WinDivert.dll")
}

// Used to call WinDivert's functions
type WinDivertHandle struct {
	handle uintptr
	open   bool
}

// LoadDLL loads the WinDivert DLL depending the OS (x64 or x86) and the given DLL path.
// The path can be a relative path (from the .exe folder) or absolute path.
func LoadDLL(path64, path32 string) {
	var dllPath string

	if runtime.GOARCH == "amd64" {
		dllPath = path64
	} else {
		dllPath = path32
	}

	winDivertDLL = syscall.NewLazyDLL(dllPath)

	winDivertOpen = winDivertDLL.NewProc("WinDivertOpen")
	winDivertClose = winDivertDLL.NewProc("WinDivertClose")
	winDivertRecv = winDivertDLL.NewProc("WinDivertRecv")
	winDivertRecvEx = winDivertDLL.NewProc("WinDivertRecvEx")
	winDivertSend = winDivertDLL.NewProc("WinDivertSend")
	winDivertSendEx = winDivertDLL.NewProc("WinDivertSendEx")
	winDivertHelperCalcChecksums = winDivertDLL.NewProc("WinDivertHelperCalcChecksums")
	winDivertHelperEvalFilter = winDivertDLL.NewProc("WinDivertHelperEvalFilter")
	winDivertHelperCheckFilter = winDivertDLL.NewProc("WinDivertHelperCheckFilter")
	winDivertGetParam = winDivertDLL.NewProc("WinDivertGetParam")
	winDivertSetParam = winDivertDLL.NewProc("WinDivertSetParam")
	winDivertHelperParsePacket = winDivertDLL.NewProc("WinDivertHelperParsePacket")
}

// Create a new WinDivertHandle by calling WinDivertOpen and returns it
// The string parameter is the fiter that packets have to match
// https://reqrypt.org/windivert-doc.html#divert_open
func NewWinDivertHandle(filter string) (*WinDivertHandle, error) {
	err := winDivertOpen.Find()
	if err != nil {
		return nil, err
	}
	return NewWinDivertHandleWithFlags(filter, 0)
}

// Create a new WinDivertHandle by calling WinDivertOpen and returns it
// The string parameter is the fiter that packets have to match
// and flags are the used flags used
// https://reqrypt.org/windivert-doc.html#divert_open
func NewWinDivertHandleWithFlags(filter string, flags uint8) (*WinDivertHandle, error) {
	filterBytePtr, err := syscall.BytePtrFromString(filter)
	if err != nil {
		return nil, err
	}

	handle, _, err := winDivertOpen.Call(uintptr(unsafe.Pointer(filterBytePtr)),
		uintptr(0),
		uintptr(0),
		uintptr(flags))

	if handle == uintptr(syscall.InvalidHandle) {
		return nil, err
	}

	winDivertHandle := &WinDivertHandle{
		handle: handle,
		open:   true,
	}
	return winDivertHandle, nil
}

// Close the Handle
// See https://reqrypt.org/windivert-doc.html#divert_close
func (wd *WinDivertHandle) Close() error {
	_, _, err := winDivertClose.Call(wd.handle)
	wd.open = false
	return err
}

// Divert a packet from the Network Stack
// https://reqrypt.org/windivert-doc.html#divert_recv
func (wd *WinDivertHandle) Recv() (*Packet, error) {
	if !wd.open {
		return nil, errors.New("can't receive, the handle isn't open")
	}

	packetBuffer := make([]byte, MaxPacketBufferSize)

	var packetLen uint
	var addr Address
	success, _, err := winDivertRecv.Call(wd.handle,
		uintptr(unsafe.Pointer(&packetBuffer[0])),
		uintptr(MaxPacketBufferSize),
		uintptr(unsafe.Pointer(&packetLen)),
		uintptr(unsafe.Pointer(&addr)))

	if success == 0 {
		return nil, err
	}

	packet := &Packet{
		Raw:       packetBuffer[:packetLen],
		Addr:      &addr,
		PacketLen: packetLen,
	}

	return packet, nil
}

// Divert a packet from the Network Stack
// https://reqrypt.org/windivert-doc.html#divert_recv_ex
func (wd *WinDivertHandle) RecvEx() ([]byte, []Address, uint, error) {
	if !wd.open {
		return nil, nil, 0, errors.New("can't receiveEx, the handle isn't open")
	}
	quantity := uint(MaxPacketBufferSize / PacketBufferSize)
	packetBuffer := make([]byte, MaxPacketBufferSize)
	var packetLen uint
	addr := make([]Address, quantity)
	size := uint(unsafe.Sizeof(Address{}))
	addrlen := size * quantity
	success, _, err := winDivertRecvEx.Call(wd.handle,
		uintptr(unsafe.Pointer(&packetBuffer[0])),
		uintptr(MaxPacketBufferSize),
		uintptr(unsafe.Pointer(&packetLen)),
		uintptr(0),
		uintptr(unsafe.Pointer(&addr[0])),
		uintptr(unsafe.Pointer(&addrlen)),
		uintptr(0))

	if success == 0 {
		return nil, nil, 0, err
	}
	packets := packetBuffer[:packetLen]
	len := addrlen / size
	addrs := addr[:len]
	return packets, addrs, len, nil
}

// Inject the packet on the Network Stack
// https://reqrypt.org/windivert-doc.html#divert_send
func (wd *WinDivertHandle) Send(packet *Packet) (uint, error) {
	var sendLen uint

	if !wd.open {
		return 0, errors.New("can't Send, the handle isn't open")
	}

	success, _, err := winDivertSend.Call(wd.handle,
		uintptr(unsafe.Pointer(&(packet.Raw[0]))),
		uintptr(packet.PacketLen),
		uintptr(unsafe.Pointer(&sendLen)),
		uintptr(unsafe.Pointer(packet.Addr)))

	if success == 0 {
		return 0, err
	}

	return sendLen, nil
}

// Inject the mutipacket on the Network Stack
// https://reqrypt.org/windivert-doc.html#divert_send_ex
func (wd *WinDivertHandle) SendEx(packets []byte, addr []Address) (uint, error) {
	var sendLen uint

	if !wd.open {
		return 0, errors.New("can't Send, the handle isn't open")
	}
	l := unsafe.Sizeof(Address{}) * uintptr(len(addr))
	success, _, err := winDivertSendEx.Call(wd.handle,
		uintptr(unsafe.Pointer(&packets[0])),
		uintptr(len(packets)),
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&addr[0])),
		uintptr(l),
		uintptr(0))

	if success == 0 {
		return 0, err
	}
	return sendLen, nil
}

// Calls WinDivertHelperCalcChecksum to calculate the packet's chacksum
// https://reqrypt.org/windivert-doc.html#divert_helper_calc_checksums
func (wd *WinDivertHandle) HelperCalcChecksum(packet *Packet) {
	winDivertHelperCalcChecksums.Call(
		uintptr(unsafe.Pointer(&packet.Raw[0])),
		uintptr(packet.PacketLen),
		uintptr(unsafe.Pointer(&packet.Addr)),
		uintptr(0))
}

// Take the given filter and check if it contains any error
// https://reqrypt.org/windivert-doc.html#divert_helper_check_filter
func HelperCheckFilter(filter string) (bool, int) {
	var errorPos uint

	filterBytePtr, _ := syscall.BytePtrFromString(filter)

	success, _, _ := winDivertHelperCheckFilter.Call(
		uintptr(unsafe.Pointer(filterBytePtr)),
		uintptr(0),
		uintptr(0), // Not implemented yet
		uintptr(unsafe.Pointer(&errorPos)))

	if success == 1 {
		return true, -1
	}
	return false, int(errorPos)
}

// GetParam
// https://reqrypt.org/windivert-doc.html#divert_get_param
func (wd *WinDivertHandle) GetParam(param WinDivertParam) (uint64, error) {
	var value uint64
	success, _, err := winDivertGetParam.Call(
		wd.handle,
		uintptr(param),
		uintptr(unsafe.Pointer(&value)))

	if success == 0 {
		return 0, err
	}

	return value, nil
}

// SetParam
// https://reqrypt.org/windivert-doc.html#divert_set_param
func (wd *WinDivertHandle) SetParam(param WinDivertParam, value uint64) error {
	success, _, err := winDivertSetParam.Call(
		wd.handle,
		uintptr(param),
		uintptr(value))

	if success == 0 {
		return err
	}

	return nil
}

// HelperParsePacket
// https://reqrypt.org/windivert-doc.html#divert_helper_parse_packet
func (wd *WinDivertHandle) HelperParsePacket(p []byte) (PayLoad, error) {
	packlen := len(p)
	packetpointer := make(PayLoad, MaxPacketBufferSize)
	var len uint
	success, _, err := winDivertHelperParsePacket.Call(
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(packlen),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&packetpointer)), //uintptr(0), uintptr(unsafe.Pointer(&packetpointer)),
		uintptr(unsafe.Pointer(&len)),
		uintptr(0), //uintptr(unsafe.Pointer(&nextpacketpointer)),
		uintptr(0), //uintptr(unsafe.Pointer(&packlen)),
	)
	if success == 0 {
		return nil, err
	}
	return packetpointer[:len], nil
}

// Take a packet and compare it with the given filter
// Returns true if the packet matches the filter
// https://reqrypt.org/windivert-doc.html#divert_helper_eval_filter
func HelperEvalFilter(packet *Packet, filter string) (bool, error) {
	filterBytePtr, err := syscall.BytePtrFromString(filter)
	if err != nil {
		return false, err
	}

	success, _, err := winDivertHelperEvalFilter.Call(
		uintptr(unsafe.Pointer(filterBytePtr)),
		uintptr(0),
		uintptr(unsafe.Pointer(&packet.Raw[0])),
		uintptr(packet.PacketLen),
		uintptr(unsafe.Pointer(&packet.Addr)))

	if success == 0 {
		return false, err
	}

	return true, nil
}

// A loop that capture packets by calling Recv and sends them on a channel as long as the handle is open
// If Recv() returns an error, the loop is stopped and the channel is closed
func (wd *WinDivertHandle) recvLoop(packetChan chan<- *Packet) {
	for wd.open {
		packet, err := wd.Recv()
		if err != nil {
			continue
			//close(packetChan)
			//break
		}

		packetChan <- packet
	}
}

func (wd *WinDivertHandle) recvLoopEx(packetChan chan<- *Packet) {
	for wd.open {
		bytes, addr, count, err := wd.RecvEx()
		if err != nil {
			continue
		}
		switch count {
		case 1:
			packet := &Packet{
				Raw:       bytes,
				Addr:      &addr[0],
				PacketLen: uint(len(bytes)),
			}
			packetChan <- packet
		default:
			for _, add := range addr {
				l, err := caculateLen(bytes)
				if err != nil {
					continue
				}
				address := add
				packet := &Packet{
					Raw:       bytes[:l],
					Addr:      &address,
					PacketLen: uint(l),
				}
				packetChan <- packet
				bytes = bytes[l:]
			}
		}

	}
}

// Create a new channel that will be used to pass captured packets and returns it calls recvLoop to maintain a loop
func (wd *WinDivertHandle) Packets() (chan *Packet, error) {
	if !wd.open {
		return nil, errors.New("the handle isn't open")
	}
	packetChan := make(chan *Packet, PacketChanCapacity)
	go wd.recvLoop(packetChan)
	return packetChan, nil
}

func (wd *WinDivertHandle) PacketExs() (chan *Packet, error) {
	if !wd.open {
		return nil, errors.New("the handle isn't open")
	}
	packetChan := make(chan *Packet, PacketChanCapacity)
	go wd.recvLoopEx(packetChan)
	return packetChan, nil
}

func caculateLen(raw []byte) (uint16, error) {
	ipVersion := int(raw[0] >> 4)
	switch ipVersion {
	case 4:
		return binary.BigEndian.Uint16(raw[2:4]), nil
	case 6:
		return binary.BigEndian.Uint16(raw[4:6]) + 40, nil
	}
	return 0, errors.New("error packet")
}
