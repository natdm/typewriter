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
	go build && ./typewriter -dir ./examples -lang flow -out ./models.js

elm:
	$(GOBUILD) -v ./
	./typewriter -dir ./stubs -lang elm

ts:
	$(GOBUILD) -v ./
	./typewriter -dir ./stubs -lang ts

testts:
	$(GOBUILD) -v ./
	./typewriter -file=./stubs/struct.go -r=false -out=./stubs/typescript -lang=typescript