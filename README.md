# segment-proxy

Proxies requests to the Segment CDN and Tracking API.

# CLI

```
Usage of proxy:
  -debug
        debug mode
  -port string
        bind address (default "8080")
```

# Usage

### Via Source

1. Clone the repo `git clone git@github.com:segmentio/segment-proxy.git`.
2. Run `make build run`.

### Via Golang

1. Run `go get github.com/segmentio/segment-proxy`.
2. Run `segment-proxy`.

### Via prebuilt binaries.

1. Run `wget https://github.com/segmentio/segment-proxy/releases/download/1.0.0/proxy`.
2. Run `chmod +x proxy`.
3. Run `./proxy`.

### Via Docker

1. Run `docker run segment/proxy`.
