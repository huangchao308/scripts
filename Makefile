# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: build
build:
	$(GOBUILD) -o bin/scripts -v ./main.go
