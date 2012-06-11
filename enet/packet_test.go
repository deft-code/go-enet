package enet

import "testing"

func TestCreatePacket(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	packet := CreatePacket(data, UNSEQUENCED)

	if int(packet.dataLength) != len(data) {
		t.Fail()
	} else {
		packet.Destroy()
	}
}
