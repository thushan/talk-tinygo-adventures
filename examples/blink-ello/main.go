package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		led.High()
		time.Sleep(time.Millisecond * 1000)
		for i := 0; i < 5; i++ {
			led.Low()
			time.Sleep(time.Millisecond * 250)

			led.High()
			time.Sleep(time.Millisecond * 250)
		}
	}
}