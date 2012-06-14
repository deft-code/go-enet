package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "errors"

func Initialize() error {
	if C.enet_initialize() != C.int(0) {
		return errors.New("ENet failed to initialize")
	}
	return nil
}

func Deinitialize() {
	C.enet_deinitialize()
}
