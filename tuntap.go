package tuntap

import (
	"os"
)

type TunTap struct {
	file *os.File
	name string
}

func (t *TunTap) Name() string {
	return t.name
}

func (t *TunTap) Read(b []byte) (int, error) {
	return t.file.Read(b)
}

func (t *TunTap) Write(b []byte) (int, error) {
	return t.file.Write(b)
}

func (t *TunTap) Close() error {
	return t.file.Close()
}
