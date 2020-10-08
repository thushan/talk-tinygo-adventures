package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hcsr04"
)

func main() {
	sensor := hcsr04.New(machine.D11, machine.D12)
	sensor.Configure()

	for {
		println("Distance: ", sensor.ReadDistance(), "mm")
		println("Pulse: ", sensor.ReadPulse(), "ms")
		time.Sleep(time.Millisecond * 250)
	}
}
