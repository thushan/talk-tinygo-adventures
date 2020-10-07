
# aliases
.PHONY: clean flash-arduino-blink-ello build-arduino-blink-ello build-tinygo-comparison build-arduino-blink-bling flash-arduino-blink-bling

clean:
	mkdir -p build
	rm -rf build/**

build-tinygo-comparison:
	go build -o build/hello-voyager_golang go-compare/main.go
	tinygo build -o build/hello-voyager_tinygo go-compare/main.go
	ls build/hello-voyager* -lah

flash-arduino-blink-ello:
	tinygo flash -target arduino -size short blink-ello/main.go

build-arduino-blink-ello:
	tinygo build -target arduino -o build/blink-ello.hex blink-ello/main.go

flash-arduino-blink-bling:
	tinygo flash -target arduino -size short -scheduler coroutines blink-bling/main.go

build-arduino-blink-bling:
	tinygo build -target arduino -o build/blink-bling.hex blink-bling/main.go