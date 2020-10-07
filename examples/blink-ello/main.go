package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ledSwitch := true
	print("Starting to blink...")

	for {

		led.Set(ledSwitch)
		ledSwitch = !ledSwitch

		println("Blinking for Voyager!")

		time.Sleep(time.Millisecond * 500)
	}
}
