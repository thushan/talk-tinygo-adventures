package main

import (
	"machine"
	"time"
)

func main() {
	go blinkOnboardLed()
	blinkLed()
}

func blinkLed() {
	led := machine.D12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Low()
		time.Sleep(1 * time.Second)

		led.High()
		time.Sleep(1 * time.Second)

		for i := 0; i < 5; i++ {
			led.Low()
			time.Sleep(250 * time.Millisecond)

			led.High()
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func blinkOnboardLed() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ledSwitch := true

	for {

		led.Set(ledSwitch)
		ledSwitch = !ledSwitch

		time.Sleep(500 * time.Millisecond)
	}
}
