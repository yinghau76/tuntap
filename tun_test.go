package tuntap

import "testing"

func TestTun(t *testing.T) {
	tun, err := NewTun("")
	if err != nil {
		t.Log("Failed to NewTun:", err)
		t.FailNow()
	}

	err = tun.Close()
	if err != nil {
		t.Log("Failed to close Tun:", err)
		t.FailNow()
	}
}
