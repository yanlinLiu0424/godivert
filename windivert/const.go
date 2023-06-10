package windivert

type Direction bool
type PayLoad []byte

const (
	PacketBufferSize   = 1500
	PacketChanCapacity = 256

	WinDivertDirectionOutbound Direction = false
	WinDivertDirectionInbound  Direction = true
)

const (
	WinDivertFlagSniff uint8 = 1 << iota
	WinDivertFlagDrop  uint8 = 1 << iota
	WinDivertFlagDebug uint8 = 1 << iota
)

type WinDivertParam uint

const (
	WinDivertParamQueueLength WinDivertParam = iota
	WinDivertParamQueueTime
	WinDivertParamQueueSize
)
const (
	WinDivertParamQueueSizeMax     = 33554432 /* 32MB */
	WinDivertParamQueueSizeDefault = 4194304  /* 4MB */
	WinDivertParamQueueSizeMin     = 65535    /* 64KB */

)

const (
	WinDivertParamQueueLengthDefault = 4096
	WinDivertParamQueueLengthMax     = 16384
	WinDivertParamQueueLengthMin     = 32
)

const (
	WinDivertParamQueueTimeDefault = 2000  /* 2s */
	WinDivertParamQueueTimeMax     = 16000 /* 16s */
	WinDivertParamQueueTimeMin     = 100   /* 100ms */
)

func (d Direction) String() string {
	if bool(d) {
		return "Inbound"
	}
	return "Outbound"
}
