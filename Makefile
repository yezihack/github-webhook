#!/usr/bin/env bash

name=go-github-webhook
VERSION=1.0

dev:run

fmt:
	gofmt -l -w ./

build:clean fmt
	go build -o bin/test/${name} .

clean:
	rm -rf bin/test/*

run:build
	bin/test/${name}



linux:fmt
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=linux
	go build -o bin/linux/${name} .

window:fmt
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=windows
	go build -o bin/window/${name}.exe .

mac:fmt
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=darwin
	go build -o bin/mac/${name} .

clear:
	rm -rf bin/window
	rm -rf bin/linux
	rm -rf bin/mac

tar:clear linux window mac
	tar -czf bin/${name}$(VERSION).window-amd64.tar.gz bin/window/${name}.exe
	tar -czf bin/${name}$(VERSION).linux-amd64.tar.gz bin/linux/${name}
	tar -czf bin/${name}$(VERSION).darwin-amd64.tar.gz bin/mac/${name}


.PHONY:fmt build clean run dev window linux mac clear tar
