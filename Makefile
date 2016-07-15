build:
	go build -o bin/segment-proxy

server:
	go run main.go

test:
	go test -v -cover ./...

docker:
	docker build -t segment/proxy .

docker-push:
	docker push segment/proxy

.PHONY: build server test docker docker-push
