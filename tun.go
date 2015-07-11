package tuntap

import (
	"os"
)

type Tun struct {
  file *os.File
  name string
}


