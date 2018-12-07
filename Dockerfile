FROM golang:alpine

COPY . $GOPATH/src/github.com/heysphere/segment-proxy
WORKDIR $GOPATH/src/github.com/heysphere/segment-proxy

RUN apk add --update make
RUN apk add --update git
RUN go get github.com/tools/godep

RUN godep restore
RUN make build

EXPOSE 8080
CMD bin/segment-proxy