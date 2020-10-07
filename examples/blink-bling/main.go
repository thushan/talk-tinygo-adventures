package main

import (
	"machine"
	"time"
)

func main() {
	go onboardled()
	led := machine.D12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledSwitch := true

	for {
		led.Set(ledSwitch)
		ledSwitch = !ledSwitch
		time.Sleep(time.Millisecond * 500)
	}

}
func onboardled() {
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
