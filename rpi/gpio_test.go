package rpi

import (
	"github.com/whaly/rpigpio"
)

// assert that rpi.pin implements gpio.Pin
var _ gpio.Pin = new(pin)
