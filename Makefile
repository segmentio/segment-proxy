build:
	gox -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"

server:
	go run main.go

test:
	go test -v -cover ./...

docker-build-test:
	docker build -t gcr.io/togather-test1/segment-proxy .

docker-build-prod:
	docker build -t gcr.io/airpict/segment-proxy .

docker-push-test:
	docker push gcr.io/togather-test1/segment-proxy

docker-push-prod:
	docker push gcr.io/airpict/segment-proxy


.PHONY: build server test docker-build-test docker-build-prod docker-push-test docker-push-prod
