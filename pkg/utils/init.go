package utils

import (
	"time"

	"GpioController/types"

	"github.com/stianeikeland/go-rpio"
)

func InitControllerFields(c *types.Controller, pin int, tickerTime int) {
	c.Pin = rpio.Pin(pin)
	c.Ticker = time.Duration(tickerTime) * time.Second
}
