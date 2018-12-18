# Go command
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
BINARY_DIR=bin

clean:
	go clean

fmt:
	gofmt -w $(GOFMT_FILES)

test: clean
	go test -v ./...

.PHONY: clean
