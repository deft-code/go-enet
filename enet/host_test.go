package enet

import "testing"
import "net"

func TestConvAddr(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "localhost:9998")

	if err != nil {
		t.Error(err)
	}

	if addr.Port != 9998 {
		t.Error(addr)
	}

	caddr := conv_addr(addr)
	if caddr.port != 9998 {
		t.Error(caddr)
	}

	cbuf := [4]byte{}
	cbuf[0] = byte(caddr.host & 0xff)
	cbuf[1] = byte((caddr.host >> 8) & 0xff)
	cbuf[2] = byte((caddr.host >> 16) & 0xff)
	cbuf[3] = byte((caddr.host >> 24) & 0xff)

	buf := [...]byte{127, 0, 0, 1}

	if cbuf != buf {
		t.Errorf("%v != %v", cbuf, buf)
	}
}

func TestCreatHost(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "localhost:9998")
	if err != nil {
		t.Error(err)
	}

	host, err := CreateHost(addr, 2, 2, 0, 0)
	if err != nil {
		t.Error(err)
	} else {
		defer host.Destroy()
	}

	if host.host.peerCount != 2 {
		t.Error(host.host.peerCount)
	}
}
