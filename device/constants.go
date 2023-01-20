package device

// OpCodes
const (
	OP_HELLO      = 0
	OP_OS_SYSTEM  = 1
	OP_SET_VOLUME = 2
	OP_GET_VOLUME = 3
	OP_SET_MUTE   = 4
	OP_GET_MUTE   = 5
)

var (
	ConnList = map[string]*WsDevice{}
)
