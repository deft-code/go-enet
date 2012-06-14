package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "unsafe"

type Flag uint

const (
	RELIABLE            Flag = C.ENET_PACKET_FLAG_RELIABLE
	UNSEQUENCED         Flag = C.ENET_PACKET_FLAG_UNSEQUENCED
	NO_ALLOCATE         Flag = C.ENET_PACKET_FLAG_NO_ALLOCATE
	UNRELIABLE_FRAGMENT Flag = C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENT
)

func new_packet(data []byte, flags Flag) *C.ENetPacket {
	packet := C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(len(data)), C.enet_uint32(flags))
	if packet == nil {
		panic("Allocation failure inside ENet")
	}
	return packet
}

func from_packet(cpacket *C.ENetPacket) []byte {
	return C.GoBytes(unsafe.Pointer(cpacket.data), C.int(cpacket.dataLength))
}
