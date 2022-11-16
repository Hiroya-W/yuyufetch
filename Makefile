NAME := yuyufetch
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

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
