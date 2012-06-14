package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

type Type uint

const (
	NONE       Type = C.ENET_EVENT_TYPE_NONE
	CONNECT    Type = C.ENET_EVENT_TYPE_CONNECT
	DISCONNECT Type = C.ENET_EVENT_TYPE_DISCONNECT
	RECEIVE    Type = C.ENET_EVENT_TYPE_RECEIVE
)

type Event struct {
	Type      Type
	Peer      Peer
	ChannelID uint8
	Data      uint32
	Packet    []byte
}

func create_event(c_event *C.ENetEvent) Event {
	return Event{
		Type:      Type(c_event._type),
		Peer:      Peer{c_event.peer},
		ChannelID: uint8(c_event.channelID),
		Data:      uint32(c_event.data),
		Packet:    from_packet(c_event.packet),
	}
}
