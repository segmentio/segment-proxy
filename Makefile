build:
	godep go build -o bin/segment-proxy

test:
	godep go test -v -cover

docker:
	docker build -t segment-proxy .
	docker run --publish 6060:8080 --name segment-proxy --rm segment-proxy

.PHONY: build test docker
