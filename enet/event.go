package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "log"

type Type uint

const (
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
		log.Fatal(args...)
	}
}

func (host *Host) create_event(c_event *C.ENetEvent) *Event {
	peer := host.peers[c_event.peer]
	e_type := Type(c_event._type)
	switch e_type {
	case CONNECT:
      log.Printf("connect: %#v",c_event)
      if peer == nil {
         peer = &Peer{c_event.peer}
         host.peers[c_event.peer] = peer
      }
	case DISCONNECT:
      log.Printf("disconnect: %#v",c_event)
		enforce(peer != nil, host, peer)
		delete(host.peers, c_event.peer)
      peer.peer = nil
	case RECEIVE:
      log.Printf("receive: %#v",c_event)
		enforce(peer != nil, host, peer)
		enforce(c_event.packet != nil)
	case Type(C.ENET_EVENT_TYPE_NONE):
		log.Printf("Discarding none event: %#v", c_event)
		return nil
	}
	return &Event{
		Type:      e_type,
		Peer:      peer,
		ChannelID: uint8(c_event.channelID),
		Data:      uint32(c_event.data),
		Packet:    from_packet(c_event.packet),
	}
}
