FROM golang:1.8
MAINTAINER Hantao Wang

EXPOSE 8082

RUN mkdir -p /go/src/github.com/kelda-inc
RUN mkdir -p /go/bin

COPY . /go/src/github.com/kelda-inc/hotrod-driver

WORKDIR /go/src/github.com/kelda-inc/hotrod-driver

RUN go build -o hotrod main.go
RUN mv hotrod /go/bin/

ENTRYPOINT ["/go/bin/hotrod", "driver"]
