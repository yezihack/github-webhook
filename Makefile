#!/usr/bin/env bash

name=github-webhook
VERSION=1.5.0

dev:run

fmt:
	gofmt -l -w ./

build:clean fmt
	go build -ldflags="-s -w" -o bin/test/${name} .

clean:
	rm -rf bin/test/*

run:build
	bin/test/${name}


cover:
	go test -cover

file:
	go test -v -covermode=count -coverprofile=${OUT}

func:file
	go tool cover -func=${OUT}

html:file
	go tool cover -html=${OUT} -o ${OUT}.html

goveralls:
	${GOPATH}/bin/goveralls -coverprofile=${OUT} -service=travis-ci -repotoken ${COVERALLS_TOKEN}


linux:fmt
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/linux/${name} .

window:fmt
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags="-s -w" -o bin/window/${name}.exe .

mac:fmt
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o bin/mac/${name} .

clear:
	rm -rf bin/window
	rm -rf bin/linux
	rm -rf bin/mac
	rm -rf bin/tar

tar:clear linux window mac clear
	mkdir -p bin/tar 
	tar -czf bin/tar/${name}$(VERSION).window-amd64.tar.gz -C bin/window ${name}.exe
	tar -czf bin/tar/${name}$(VERSION).linux-amd64.tar.gz -C bin/linux ${name}
	tar -czf bin/tar/${name}$(VERSION).darwin-amd64.tar.gz -C bin/mac ${name}


.PHONY:fmt build clean run dev window linux mac clear tar
