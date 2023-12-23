.PHONY: all

build:
	go build -o ./bin/ -ldflags="-X main.AppVersion=$(shell git describe --tag --always)" github.com/246859/goplay/cmd/goplay