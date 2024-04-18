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
	@echo "\n> removing the bin file ./bin/suitetalk"
	@make clean

## clean: remove the bin file 
.PHONY: clean
clean:
	@go clean
	rm ./bin/${BINARY_NAME}

#---------------------------------------
#|       Dummy Commands for Test       |
#---------------------------------------
.PHONY: config\:account
config\:account: build
	@./bin/${BINARY_NAME} config account --name sandbox --passport 'account=23423_SB1,cosumer-secret="consumersecret",consumer-key="consumerkey",token-id="tokenid", token-secret="tokensecret"' || echo "Error en la ejecución de ./bin/${BINARY_NAME}"
	@make clean

.PHONY: create\:vendor
create\:vendor: build
	@./bin/${BINARY_NAME} create --type vendor --data '{}' || echo "Error en la ejecución de ./bin/${BINARY_NAME}"
	@make clean