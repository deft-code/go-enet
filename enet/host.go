package enet

/*
#cgo LDFLAGS: -lenet
#include <enet/enet.h>
*/
import "C"

import "errors"
import "net"
import "time"
import "unsafe"

type Host struct {
	host *C.ENetHost
}

func conv_addr(address net.UDPAddr) (ret C.ENetAddress) {
	c_ip := C.CString(address.IP.String())
	defer C.free(unsafe.Pointer(c_ip))

	C.enet_address_set_host(&ret, c_ip)
	ret.port = C.enet_uint16(address.Port)
	return
}

// enet_host_create
func CreateHost(address net.UDPAddr, peerCount uint, channelLimit uint, incomingBandwidth uint32, outgoingBandwith uint32) (Host, error) {
	addr := conv_addr(address)
	host := C.enet_host_create(&addr, C.size_t(peerCount), C.size_t(channelLimit),
		C.enet_uint32(incomingBandwidth), C.enet_uint32(outgoingBandwith))

	if host == nil {
		return Host{nil}, errors.New("ENet failed to create host")
	}
	return Host{host}, nil
}

// enet_host_destroy
func (host Host) Destory() {
	C.enet_host_destroy(host.host)
	//host.host = uintptr(0)
}

// enet_host_connect
func (host Host) Connect(address net.UDPAddr, channelCount uint, data uint) (Peer, error) {
	addr := conv_addr(address)
	peer := C.enet_host_connect(host.host, &addr, C.size_t(channelCount), C.enet_uint32(data))

	if peer == nil {
		return Peer{nil}, errors.New("ENet failed to connect")
	}
	return Peer{peer}, nil
}

// enet_host_flush
func (host Host) Flush() {
	C.enet_host_flush(host.host)
}

// enet_host_broadcast
func (host Host) Broadcast(channelID uint8, packet []byte, flags Flag) {
	C.enet_host_broadcast(host.host, C.enet_uint8(channelID), new_packet(packet, flags))
}

func ret_to_error(cevent *C.ENetEvent, ret C.int) (Event, error) {
	switch {
	case ret < 0:
		return Event{Type: NONE}, errors.New("ENet error")
	case ret == 0:
		return Event{Type: NONE}, nil
	case ret > 0:
		return create_event(cevent), nil
	}
   panic("All cases covered")
}

// enet_host_service
func (host Host) Service(timeout time.Duration) (Event, error) {
	if timeout < 0 {
		return Event{Type: NONE}, errors.New("Timeout duration was negative")
	}

	var cevent C.ENetEvent

	ret := C.enet_host_service(host.host, &cevent, C.enet_uint32(timeout/time.Millisecond))

	return ret_to_error(&cevent, ret)
}

// enet_host_check_event
func (host Host) CheckEvents() (Event, error) {
	var cevent C.ENetEvent

	ret := C.enet_host_check_events(host.host, &cevent)

	return ret_to_error(&cevent, ret)
}
