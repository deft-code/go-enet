package enet

import "testing"

func TestCreatePacket(t *testing.T) {
	data := []data{1, 2, 3, 4, 5}
	packet := CreatePacket(data, UNRELIABLE)

	if packet.dataLength != len(data) {
		t.Fail()
	} else {
		defer packet.Destroy()
	}
}
