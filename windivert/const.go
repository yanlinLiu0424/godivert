package windivert

type Direction bool

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

type WINDIVERT_PARAM uint

const (
	WINDIVERT_PARAM_QUEUE_LENGTH WINDIVERT_PARAM = iota
	WINDIVERT_PARAM_QUEUE_TIME
	WINDIVERT_PARAM_QUEUE_SIZE
)
const (
	WINDIVERT_PARAM_QUEUE_SIZE_MAX     = 33554432 /* 32MB */
	WINDIVERT_PARAM_QUEUE_SIZE_DEFAULT = 4194304  /* 4MB */
	WINDIVERT_PARAM_QUEUE_SIZE_MIN     = 65535    /* 64KB */

)

const (
	WINDIVERT_PARAM_QUEUE_LENGTH_DEFAULT = 4096
	WINDIVERT_PARAM_QUEUE_LENGTH_MAX     = 16384
	WINDIVERT_PARAM_QUEUE_LENGTH_MIN     = 32
)

const (
	WINDIVERT_PARAM_QUEUE_TIME_DEFAULT = 2000  /* 2s */
	WINDIVERT_PARAM_QUEUE_TIME_MAX     = 16000 /* 16s */
	WINDIVERT_PARAM_QUEUE_TIME_MIN     = 100   /* 100ms */
)

func (d Direction) String() string {
	if bool(d) {
		return "Inbound"
	}
	return "Outbound"
}
