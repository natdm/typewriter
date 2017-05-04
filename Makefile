GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test

.PHONY: test

test:
	$(GOTEST) -v ./template

flow:
	$(GOBUILD) -v ./
	go build && ./typewriter -file ./examples/example.go -lang flow -out ./models.js -v

elm:
	$(GOBUILD) -v ./
	go build && ./typewriter -dir ./examples -lang elm -out ./models.js

ts:
	$(GOBUILD) -v ./
	go build && ./typewriter -dir ./examples -lang ts -out ./models.js

testts:
	$(GOBUILD) -v ./
	./typewriter -file=./stubs/struct.go -r=false -out=./stubs/typescript -lang=typescript