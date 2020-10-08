
# aliases
.PHONY: clean build-tinygo-comparison flash-arduino-blink-ello build-arduino-blink-ello build-arduino-blink-bling flash-arduino-blink-bling flash-arduino-nano-blink-ello build-arduino-nano-blink-ello build-arduino-nano-blink-bling flash-arduino-nano-blink-bling build-arduino-temp-sensor flash-arduino-temp-sensor build-arduino-distance-sensor flash-arduino-distance-sensor

clean:
	mkdir -p _build
	rm -rf _build/**

build-tinygo-comparison:
	go build -o _build/hello-voyager_golang go-compare/main.go
	tinygo build -o _build/hello-voyager_tinygo go-compare/main.go
	go version && tinygo version && ls _build/hello-voyager* -lah	

flash-arduino-nano-blink-ello:
	tinygo flash -target arduino-nano -size short blink-ello/main.go

build-arduino-nano-blink-ello:
	tinygo build -target arduino-nano -o _build/blink-ello.hex blink-ello/main.go

flash-arduino-nano-blink-bling:
	tinygo flash -target arduino-nano -size short -scheduler coroutines blink-bling/main.go

build-arduino-nano-blink-bling:
	tinygo build -target arduino-nano -o _build/blink-bling.hex blink-bling/main.go

flash-arduino-blink-ello:
	tinygo flash -target arduino -size short blink-ello/main.go

build-arduino-blink-ello:
	tinygo build -target arduino -o _build/blink-ello.hex blink-ello/main.go

flash-arduino-blink-bling:
	tinygo flash -target arduino -size short -scheduler coroutines blink-bling/main.go

build-arduino-blink-bling:
	tinygo build -target arduino -o _build/blink-bling.hex blink-bling/main.go

build-arduino-temp-sensor:
	tinygo build -target arduino -o _build/temp-sensor.hex temp-sensor/main.go

flash-arduino-temp-sensor:
	tinygo flash -target arduino -size short temp-sensor/main.go

build-arduino-distance-sensor:
	tinygo build -target arduino -o _build/distance-sensor.hex distance-sensor/main.go

flash-arduino-distance-sensor:
	tinygo flash -target arduino -size short distance-sensor/main.go