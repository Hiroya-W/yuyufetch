NAME := yuyufetch
VERSION := $(shell git describe --tags)
LDFLAGS := "-X github.com/Hiroya-W/yuyufetch.version=$(VERSION)"

## build binaries
bin/%: cmd/%/main.go
	go build -ldflags $(LDFLAGS) -o $@ $<

## build binary
.PHONY: build
build: bin/$(NAME)

## clean
.PHONY: clean
clean:
	rm ./bin/$(NAME)
