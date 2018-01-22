# This is the dockerfile we use to run the project locally as well as
# compile the code for a slim production image
FROM golang:1.8.3-alpine3.6

ENV CGO_ENABLED=0\
    GOOS=linux

WORKDIR /go/src/github.com/InVisionApp/sio-proxy

# Add rest of source code
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o segment-proxy .
RUN cp segment-proxy /segment-proxy

ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/segment-proxy"]
