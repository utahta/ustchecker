.PHONY: install build fmt test

all: build

install:
	@glide install

build:
	@go build ./cmd/uststat

fmt:
	gofmt -w .
	goimports -w .

test:
	@go test -v -race 

