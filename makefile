.PHONY: all

build:
	go build -o ./bin/play.exe -ldflags="-X main.AppVersion=$(shell git describe --tag --always)" github.com/246859/goplay/cmd