package dev

import (
	"github.com/LassiHeikkila/ble"
	"github.com/LassiHeikkila/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
