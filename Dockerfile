FROM golang:1.8
MAINTAINER Hantao Wang

EXPOSE 8082

RUN mkdir -p /go/src/gitlab.com/will.wang1
RUN mkdir -p /go/bin

WORKDIR /go/src/gitlab.com/will.wang1

ARG git_pass
ARG build_time=1

RUN git clone https://user:$git_pass@gitlab.com/will.wang1/hotrod-driver

WORKDIR /go/src/gitlab.com/will.wang1/hotrod-driver

RUN go build -o hotrod main.go
RUN mv hotrod /go/bin/

ENTRYPOINT ["/go/bin/hotrod", "driver"]
