package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>\
*/
import "C"

type Peer struct {
   peer *C.ENetPeer
}

// enet_peer_send
func (peer Peer) Send(channelID uint8, data []byte, flags uint) error {
   packet := new_packet( data, flags )
   defer C.enet_packet_destroy(packet)

   return zero_or_error( C.enet_peer_send(
         peer.peer, C.enet_uint8(channelID), packet))
}

// enet_peer_receive
func (peer Peer) Receive() ([]byte,uint8)  {
	var channel_id C.enet_uint8 = 0
	packet := C.enet_peer_receive(peer.peer, &channel_id)
	return from_packet(packet), uint8(channel_id)
}

// enet_peer_reset
func (peer Peer) Reset() {
	C.enet_peer_reset(peer.peer)
}

// enet_peer_ping
func (peer Peer) Ping() {
	C.enet_peer_ping(peer.peer)
}

// enet_peer_disconnect_now
func (peer Peer) DisconnectNow(data uint) {
	C.enet_peer_disconnect_now(peer.peer, C.enet_uint32(data))
}

// enet_peer_disconnect
func (peer Peer) Disconnect(data uint) {
	C.enet_peer_disconnect(peer.peer, C.enet_uint32(data))
}

// enet_peer_disconnect_later
func (peer Peer) DisconnectLater(data uint) {
	C.enet_peer_disconnect_later(peer.peer, C.enet_uint32(data))
}
