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
	host  *C.ENetHost
	peers map[*C.ENetPeer]*Peer
}

func conv_addr(address *net.UDPAddr) (c_addr C.ENetAddress) {
	c_ip := C.CString(address.IP.String())
	defer C.free(unsafe.Pointer(c_ip))

	C.enet_address_set_host(&c_addr, c_ip)
	c_addr.port = C.enet_uint16(address.Port)
	return
}

// enet_host_create
func CreateHost(address *net.UDPAddr, peerCount uint, channelLimit uint, incomingBandwidth uint32, outgoingBandwith uint32) (*Host, error) {
	var c_host *C.ENetHost
	if address != nil {
		c_addr := conv_addr(address)
		c_host = C.enet_host_create(&c_addr, C.size_t(peerCount), C.size_t(channelLimit),
			C.enet_uint32(incomingBandwidth), C.enet_uint32(outgoingBandwith))
	} else {
		c_host = C.enet_host_create(nil, C.size_t(peerCount), C.size_t(channelLimit),
			C.enet_uint32(incomingBandwidth), C.enet_uint32(outgoingBandwith))
	}

	if c_host == nil {
		return nil, errors.New("ENet failed to create an ENetHost.")
	}
	return &Host{c_host, make(map[*C.ENetPeer]*Peer)}, nil
}

// enet_host_destroy
func (host *Host) Destory() {
	C.enet_host_destroy(host.host)
	//host.host = uintptr(0)
}

// enet_host_connect
func (host *Host) Connect(address *net.UDPAddr, channelCount uint, data uint) (*Peer, error) {
	c_addr := conv_addr(address)
	c_peer := C.enet_host_connect(host.host, &c_addr, C.size_t(channelCount), C.enet_uint32(data))

	if c_peer == nil {
		return nil, errors.New("No available peers for initiating an ENet connection.")
	}

	peer := &Peer{c_peer}
	host.peers[c_peer] = peer

	return peer, nil
}

// enet_host_flush
func (host *Host) Flush() {
	C.enet_host_flush(host.host)
}

// enet_host_broadcast
func (host *Host) Broadcast(channelID uint8, packet []byte, flags Flag) {
	C.enet_host_broadcast(host.host, C.enet_uint8(channelID), new_packet(packet, flags))
}

func (host *Host) ret_to_error(c_event *C.ENetEvent, ret C.int) (*Event, error) {
	switch {
	case ret < 0:
		return nil, errors.New("ENet internal error")
	case ret == 0:
		return nil, nil
	case ret > 0:
		return host.create_event(c_event), nil
	}
	panic("All cases covered")
}

// enet_host_service
func (host *Host) Service(timeout time.Duration) (*Event, error) {
	if timeout < 0 {
		return nil, errors.New("Timeout duration was negative")
	}

	var c_event C.ENetEvent

	ret := C.enet_host_service(host.host, &c_event, C.enet_uint32(timeout/time.Millisecond))

	return host.ret_to_error(&c_event, ret)
}

// enet_host_check_event
func (host *Host) CheckEvents() (*Event, error) {
	var c_event C.ENetEvent

	ret := C.enet_host_check_events(host.host, &c_event)

	return host.ret_to_error(&c_event, ret)
}
