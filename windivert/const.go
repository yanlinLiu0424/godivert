package windivert

type Direction bool
type PayLoad []byte

const (
	PacketBufferSize                     = 1500
	PacketChanCapacity                   = 256
	MaxPacketBufferSize                  = 65535
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

type Layer int

const (
	LayerNetwork        Layer = 0
	LayerNetworkForward Layer = 1
	LayerFlow           Layer = 2
	LayerSocket         Layer = 3
	LayerReflect        Layer = 4
	LayerEthernet       Layer = 5
)

func (l Layer) String() string {
	switch l {
	case LayerNetwork:
		return "WINDIVERT_LAYER_NETWORK"
	case LayerNetworkForward:
		return "WINDIVERT_LAYER_NETWORK_FORWARD"
	case LayerFlow:
		return "WINDIVERT_LAYER_FLOW"
	case LayerSocket:
		return "WINDIVERT_LAYER_SOCKET"
	case LayerReflect:
		return "WINDIVERT_LAYER_REFLECT"
	case LayerEthernet:
		return "WINDIVERT_LAYER_ETHERNET"
	default:
		return ""
	}
}

type Event int

const (
	EventNetworkPacket   Event = 0
	EventFlowEstablished Event = 1
	EventFlowDeleted     Event = 2
	EventSocketBind      Event = 3
	EventSocketConnect   Event = 4
	EventSocketListen    Event = 5
	EventSocketAccept    Event = 6
	EventSocketClose     Event = 7
	EventReflectOpen     Event = 8
	EventReflectClose    Event = 9
	EventEthernetFrame   Event = 10
)

func (e Event) String() string {
	switch e {
	case EventNetworkPacket:
		return "WINDIVERT_EVENT_NETWORK_PACKET"
	case EventFlowEstablished:
		return "WINDIVERT_EVENT_FLOW_ESTABLISHED"
	case EventFlowDeleted:
		return "WINDIVERT_EVENT_FLOW_DELETED"
	case EventSocketBind:
		return "WINDIVERT_EVENT_SOCKET_BIND"
	case EventSocketConnect:
		return "WINDIVERT_EVENT_SOCKET_CONNECT"
	case EventSocketListen:
		return "WINDIVERT_EVENT_SOCKET_LISTEN"
	case EventSocketAccept:
		return "WINDIVERT_EVENT_SOCKET_ACCEPT"
	case EventSocketClose:
		return "WINDIVERT_EVENT_SOCKET_CLOSE"
	case EventReflectOpen:
		return "WINDIVERT_EVENT_REFLECT_OPEN"
	case EventReflectClose:
		return "WINDIVERT_EVENT_REFLECT_CLOSE"
	case EventEthernetFrame:
		return "WINDIVERT_EVENT_ETHERNET_FRAME"
	default:
		return ""
	}
}

type Shutdown int

const (
	ShutdownRecv Shutdown = 0
	ShutdownSend Shutdown = 1
	ShutdownBoth Shutdown = 2
)

func (h Shutdown) String() string {
	switch h {
	case ShutdownRecv:
		return "WINDIVERT_SHUTDOWN_RECV"
	case ShutdownSend:
		return "WINDIVERT_SHUTDOWN_SEND"
	case ShutdownBoth:
		return "WINDIVERT_SHUTDOWN_BOTH"
	default:
		return ""
	}
}
