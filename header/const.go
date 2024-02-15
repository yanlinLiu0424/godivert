package header

const (
	IPv4HeaderLen    = 20
	MaxIPv4HeaderLen = 60
	IPv6HeaderLen    = 40
	IPvExtensionLen  = 8
	TCPHeaderLen     = 20
	MaxTCPHeaderLen  = 60
	UDPHeaderLen     = 8
	ICMPv4HeaderLen  = 8
	ICMPv6HeaderLen  = 8

	HOPOPT            = 0
	ICMPv4            = 1
	TCP               = 6
	UDP               = 17
	ICMPv6            = 58
	IPv6Encapsulation = 41
	IPv4              = 4
	IPv6              = 6
	IPv6Route         = 43
	IPv6Frag          = 44
	ESP               = 50
	AH                = 51
	IPv6Opts          = 60
)
