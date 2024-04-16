BINARY_NAME=suitetalk

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: creates the bin file into ./bin directory
.PHONY: build
build:
	go build -o ./bin/${BINARY_NAME} main.go

## run: execute the binary field and run the root suitetalk command
.PHONY: run
run: build
	./bin/${BINARY_NAME}

## crean: remove the bin file 
.PHONY: clean
clean:
	go clean
	rm ./bin/${BINARY_NAME}