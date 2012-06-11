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
func (peer Peer) Send(channelID uint8, packet Packet) int {
	return int(C.enet_peer_send(peer, C.enet_uint8(channelID), packet))
}

// enet_peer_receive
func (peer Peer) Receive() (channelID uint8, packet []byte) {
	var c_channel_id C.enet_uint8
	packet = C.enet_peer_receive(peer, &c_channel_id, packet)
	channelID = c_channel_id
	return
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
