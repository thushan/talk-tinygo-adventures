
# aliases
.PHONY: clean flash-arduino-blink-ello build-arduino-blink-ello build-tinygo-comparison

clean:
	mkdir -p build
	rm -rf build/**

flash-arduino-blink-ello:
	tinygo flash -target arduino blink-ello/blink-ello.go

build-arduino-blink-ello:
	tinygo build -target arduino -o build/blink-ello.hex blink-ello/blink-ello.go

build-tinygo-comparison:
	go build -o build/hello-voyager_golang go-compare/go-compare.go
	tinygo build -o build/hello-voyager_tinygo go-compare/go-compare.go
	ls build/hello-voyager* -lah