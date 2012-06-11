package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "log"
import "unsafe"

const (
   RELIABLE = C.ENET_PACKET_FLAG_RELIABLE
   UNSEQUENCED = C.ENET_PACKET_FLAG_UNSEQUENCED
   NO_ALLOCATE = C.ENET_PACKET_FLAG_NO_ALLOCATE
   UNRELIABLE_FRAGMENT = C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENT
)


type ENetPacket *C.ENetPacket

// enet_packet_create
func CreatePacket(data []byte, flags int) ENetPacket {
	packet := C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(len(data)), C.enet_uint32(flags))
	if unsafe.Pointer(packet) == unsafe.Pointer(uintptr(0)) {
		log.Fatal("TODO return an error here")
	}
	return ENetPacket{packet}
}

// enet_packet_destroy
func (packet ENetPacket) Destroy() {
	C.enet_packet_destroy(packet)
}

// enet_packet_resize
func (packet ENetPacket) Resize(dataLength uint) int {
	return int(C.enet_packet_resize(packet))
}