package tuntap

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type ifreq struct {
	name    [16]byte
	flags   uint16
	padding [40 - 16 - 2]byte
}

func initTunTap(fd uintptr, name string, flags uint16) (string, error) {
	var req ifreq

	req.flags = flags
	copy(req.name[0:len(req.name)-1], name)

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&req)))
	if errno != 0 {
		return "", errno
	}
	return strings.Trim(string(req.name[:]), "\x00"), nil
}

const (
	cIFF_TUN   = 0x0001
	cIFF_TAP   = 0x0002
	cIFF_NO_PI = 0x1000

	cDEVICE_PATH = "/dev/net/tun"
)

func NewTunTap(name string, flags uint16) (*TunTap, error) {
	file, err := os.OpenFile(cDEVICE_PATH, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	devName, err := initTunTap(file.Fd(), name, flags)
	if err != nil {
		return nil, err
	}
	return &TunTap{file, devName}, nil
}

func NewTap(name string) (*Tap, error) {
	tuntap, err := NewTunTap(name, cIFF_TAP|cIFF_NO_PI)
	if err != nil {
		return nil, err
	}
	return &Tap{tuntap}, nil
}

func NewTun(name string) (*Tun, error) {
	tuntap, err := NewTunTap(name, cIFF_TUN|cIFF_NO_PI)
	if err != nil {
		return nil, err
	}
	return &Tun{tuntap}, nil
}
