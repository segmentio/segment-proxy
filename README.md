# segment-proxy

Proxies requests to the Segment CDN and Tracking API. You can run this in your infrastructure (or use it as a reference implementation) and point your clients to use this proxy instead of `api.segment.io` or `cdn.segment.com` directly.

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
2. Run `make build server`.

### Via Golang

1. Run `go get github.com/segmentio/segment-proxy`.
2. Run `segment-proxy`.

### Via prebuilt binaries.

Download the latest binaries from [Github](https://github.com/segmentio/segment-proxy/releases).

### Via Docker

1. Run `docker run --publish 8080:8080 segment/proxy`.

# Library Instructions

* [iOS](https://segment.com/docs/connections/sources/catalog/libraries/mobile/ios/#proxy-https-calls)
* [Android](https://segment.com/docs/connections/sources/catalog/libraries/mobile/android/#proxying-http-calls)
* [Analytics.js](https://segment.com/docs/connections/sources/catalog/libraries/website/javascript/custom-proxy)
