package tuntap

import (
	"testing"
)

func TestTap(t *testing.T) {
	tap, err := NewTap("")
	if err != nil {
		t.Log("Failed to NewTan:", err)
		t.FailNow()
	}

	err = tap.Close()
	if err != nil {
		t.Log("Failed to close Tun:", err)
		t.FailNow()
	}
}
