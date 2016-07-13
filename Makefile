build:
	go build -o bin/segment-proxy

test:
	go test -v -cover

.PHONY: build test
