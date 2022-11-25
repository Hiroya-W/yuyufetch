NAME := yuyufetch
VERSION := $(shell git describe --tags)
LDFLAGS := "-X github.com/Hiroya-W/yuyufetch/lib.version=$(VERSION)"

## build binaries
bin/%: main.go
	go build -ldflags $(LDFLAGS) -o $@ $<

## build binary
.PHONY: build
build: bin/$(NAME)

## clean
.PHONY: clean
clean:
	rm ./bin/$(NAME)
