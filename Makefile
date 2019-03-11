
version := $(shell git describe --tags 2> /dev/null)
buildtime := $(shell date +%Y/%m/%dT%H:%M:%S)
commit := $(shell git rev-parse --short HEAD)

build_flag := "-X main.version=develop -X main.date=$(buildtime) -X main.commit=$(commit)"

.PHONY: all

all:
	@echo "make <cmd>"

run:
	go run -v -ldflags $(build_flag)  main.go

run-service:
	go run -v -ldflags $(build_flag)  main.go service

install:
	go install -a -v -ldflags $(build_flag)

test:
	go test ./...
