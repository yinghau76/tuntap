package tuntap

import (
	"os"
)

type Tap struct {
  file *os.File
  name string
}

