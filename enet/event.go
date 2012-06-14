package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "log"

type Type uint

const (
	//NONE       Type = C.ENET_EVENT_TYPE_NONE
	CONNECT    Type = C.ENET_EVENT_TYPE_CONNECT
	DISCONNECT Type = C.ENET_EVENT_TYPE_DISCONNECT
	RECEIVE    Type = C.ENET_EVENT_TYPE_RECEIVE
)

type Event struct {
	Type      Type
	Peer      *Peer
	ChannelID uint8
	Data      uint32
	Packet    []byte
}

func enforce(condition bool, args ...interface{}) {
	if !condition {
		log.Panic(args...)
	}
}

func (host *Host) create_event(c_event *C.ENetEvent) *Event {
	peer := host.peers[c_event.peer]
	e_type := Type(c_event._type)
	switch e_type {
	case CONNECT:
		enforce(peer == nil, host, peer)
		peer = &Peer{c_event.peer}
		host.peers[c_event.peer] = peer
	case DISCONNECT:
		enforce(peer != nil, host, peer)
		delete(host.peers, c_event.peer)
	case RECEIVE:
		enforce(peer != nil, host, peer)
		enforce(c_event.packet != nil)
   default:
		enforce(false, "This should never occur", e_type)
	}
	return &Event{
		Type:      e_type,
		Peer:      peer,
		ChannelID: uint8(c_event.channelID),
		Data:      uint32(c_event.data),
		Packet:    from_packet(c_event.packet),
	}
}
