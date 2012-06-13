package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "log"
import "net"
import "unsafe"

type Host struct {
   host *C.ENetHost
}

func conv_addr(address net.UDPAddr) (ret C.ENetAddress) {
	c_ip := C.CString(address.IP.String())
   defer C.free(c_ip)

	C.enet_address_set_host(&ret, c_ip)
	ret.port = C.enet_uint16(address.Port)
	return
}

// enet_host_create
func CreateHost(address net.UDPAddr, peerCount uint, channelLimit uint,
      incomingBandwidth uint32, outgoingBandwith uint32) Host {
   addr := conv_addr(address)
   host := C.enet_host_create(&addr, C.size_t(peerCount),
      C.size_t(channelLimit),
      C.enet_uint32(incomingBandwidth),
      C.enet_uint32(outgoingBandwith))

   if unsafe.Pointer(host) == unsafe.Pointer(0) {
      log.Fatal("TODO return an error here")
   }
   return Host{host}, 
}

// enet_host_destory
func (host Host) Destory() {
	C.enet_host_destroy(host.host)
   host.host = uintptr(0)
}

// enet_host_connet
func (host Host) Connect(address net.UDPAddr, channelCount uint, data uint) ENetPeer {
	addr := conv_addr(address)
	peer := C.enet_host_connect(host, &addr, C.size_t(channelCount), C.enet_uint32(data))
	if unsafe.Pointer(peer) == unsafe.Pointer(0) {
		log.Fatal("TODO return an error here")
	}
	return peer
}

// enet_host_flush
func (host Host) Flush() {
	C.enet_host_flush(host.host)
}

// enet_host_service
func (host Host) Service(event ENetEvent, timeout_ms uint) int {
	return int(C.enet_host_service(host, event, C.enet_uint32(timeout_ms)))
}

// enet_host_broadcast
func (host Host) Broadcast(channelID uint8, packet ENetPacket) {
	C.enet_host_broadcast(host.host, C.enet_uint8(channelID), packet)
}
