package types

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

type Controller struct {
	Pin      rpio.Pin
	State    rpio.State
	Ticker   time.Duration
	PulsTime time.Duration
}
