package enet

import "testing"

func TestInitialize(t *testing.T) {
	err := Initialize()
	if err != nil {
		t.Fatal(err)
	}
	defer Deinitialize()
}
