.ONESHELL:
SHELL=/bin/bash

build:
	go build ./cmd/protoc-gen-go-json/

install:
	go install ./cmd/protoc-gen-go-json/

test: build
	protoc  --plugin=protoc-gen-go-json=./protoc-gen-go-json --go-json_out=. --go-json_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative  ./example/example.proto
	go test ./example/... 

