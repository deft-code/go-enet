package enet

import "testing"

func TestInitialize(t *testing.T) {
	if Initialize() == 0 {
		defer Deinitialize()
	} else {
		t.Fail()
	}
}
