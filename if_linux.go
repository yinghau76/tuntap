package tuntap

import (
	"os"
	"syscall"
	"unsafe"
)

type ifreq struct {
  name [16]byte
  flags uint16
  padding [40 - 16 - 2]byte
}

func createInterface(fd uintptr, name string, flags uint16) error {
  var req ifreq

  req.flags = flags
  copy(req.name[0:len(req.name)-1], name)

  _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&req)))
  return err
}

const (
  IFF_TUN = 0x0001
  IFF_TAP = 0x0002
  IFF_NO_PI = 0x1000

  DEVICE_PATH = "/dev/net/tun"
)

func NewTap(name string) (*Tap, error) {
  file, err := os.OpenFile(DEVICE_PATH, os.O_RDWR, 0)
  if err != nil {
    return nil, err
  }

  err = createInterface(file.Fd(), name, IFF_TAP | IFF_NO_PI)
  if err != nil {
    return nil, err
  }
  return &Tap{file, name}, nil
}

func NewTun(name string) (*Tun, error) {
  file, err := os.OpenFile(DEVICE_PATH, os.O_RDWR, 0)
  if err != nil {
    return nil, err
  }

  err = createInterface(file.Fd(), name, IFF_TAP | IFF_NO_PI)
  if err != nil {
    return nil, err
  }
  return &Tun{file, name}, nil
}