package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

func Initialize() int {
	return int(C.enet_initialize())
}

func Deinitialize() {
	C.enet_deinitialize()
}
