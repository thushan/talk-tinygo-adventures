
# aliases
.PHONY: clean build-tinygo-comparison flash-arduino-blink-ello build-arduino-blink-ello build-arduino-blink-bling flash-arduino-blink-bling flash-arduino-nano-blink-ello build-arduino-nano-blink-ello build-arduino-nano-blink-bling flash-arduino-nano-blink-bling

clean:
	mkdir -p _build
	rm -rf _build/**

build-tinygo-comparison:
	go build -o _build/hello-voyager_golang go-compare/main.go
	tinygo build -o _build/hello-voyager_tinygo go-compare/main.go
	ls _build/hello-voyager* -lah

flash-arduino-blink-ello:
	tinygo flash -target arduino -size short blink-ello/main.go

build-arduino-blink-ello:
	tinygo build -target arduino -o _build/blink-ello.hex blink-ello/main.go

flash-arduino-blink-bling:
	tinygo flash -target arduino -size short -scheduler coroutines blink-bling/main.go

build-arduino-blink-bling:
	tinygo build -target arduino -o _build/blink-bling.hex blink-bling/main.go

flash-arduino-nano-blink-ello:
	tinygo flash -target arduino-nano -size short blink-ello/main.go

build-arduino-nano-blink-ello:
	tinygo build -target arduino-nano -o _build/blink-ello.hex blink-ello/main.go

flash-arduino-nano-blink-bling:
	tinygo flash -target arduino-nano -size short -scheduler coroutines blink-bling/main.go

build-arduino-nano-blink-bling:
	tinygo build -target arduino-nano -o _build/blink-bling.hex blink-bling/main.go