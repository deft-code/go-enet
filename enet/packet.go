package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "unsafe"
import "log"

type Flag uint

const (
	RELIABLE            Flag = C.ENET_PACKET_FLAG_RELIABLE
	UNSEQUENCED         Flag = C.ENET_PACKET_FLAG_UNSEQUENCED
	NO_ALLOCATE         Flag = C.ENET_PACKET_FLAG_NO_ALLOCATE
	UNRELIABLE_FRAGMENT Flag = C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENT
)

func new_packet(data []byte, flags Flag) *C.ENetPacket {
	c_packet := C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(len(data)), C.enet_uint32(flags))
	if c_packet == nil {
		panic("Allocation failure inside ENet")
	}
	return c_packet
}

func from_packet(c_packet *C.ENetPacket) []byte {
   log.Printf("from packet: %#v",c_packet)
   if c_packet == nil {
      return nil
   }
	defer C.enet_packet_destroy(c_packet)
	return C.GoBytes(unsafe.Pointer(c_packet.data), C.int(c_packet.dataLength))
}
