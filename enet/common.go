package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "fmt"
import "unsafe"

const (
   RELIABLE uint = C.ENET_PACKET_FLAG_RELIABLE
   UNSEQUENCED uint = C.ENET_PACKET_FLAG_UNSEQUENCED
   NO_ALLOCATE uint = C.ENET_PACKET_FLAG_NO_ALLOCATE
   UNRELIABLE_FRAGMENT uint = C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENT
)

func is_null(ptr unsafe.Pointer) {
   return ptr == unsafe.Pointer(uintptr(0))
}

func new_packet( data []bytes, flags uint ) *C.ENetPacket {
	packet := C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(len(data)), C.enet_uint32(flags))
   enforce(!is_null(packet), "this should never happen")
	return packet
}

func from_packet( packet *C.ENetPacket ) []byte {
   ret := make([]byte, packet.dataLength)
   for i:=0; i<len(ret); i++ {
      ret[i] = byte(packet.data[i])
   }
}

func zero_or_error(val C.int) error {
   if val == 0 {
      return fmt.Errorf("ENet error: %d", ret)
   } else {
      return nil
   }
}

func enforce(condition bool, msg string) {
   if (!condition) {
      panic(msg)
   }
}
