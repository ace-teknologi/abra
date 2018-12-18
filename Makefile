# Go command
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_DIR=bin

clean:
	$(GOCLEAN)

test: clean
	$(GOTEST) -v ./...

.PHONY: clean
