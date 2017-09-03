GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test

.PHONY: test

test:
	$(GOTEST) -v ./template

build: 
	$(GOBUILD) -v ./

flow:
	./typewriter -file ./examples/example.go -lang flow -out ./models.js -v

elm:
	./typewriter -dir ./examples -lang elm -out ./models.js

ts:
	./typewriter -dir ./examples -lang ts -out ./models.js

testts:
	./typewriter -file=./stubs/struct.go -r=false -out=./stubs/typescript -lang=typescript