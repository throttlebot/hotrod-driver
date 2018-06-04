FROM golang:1.8
MAINTAINER Hantao Wang
EXPOSE 8082

RUN mkdir -p /go/src/gitlab.com/kelda-hotrod
RUN mkdir -p /go/bin
RUN go get github.com/go-redis/redis
RUN go get github.com/lib/pq
RUN go get github.com/sirupsen/logrus

WORKDIR /go/src/gitlab.com/kelda-hotrod

RUN git clone git@gitlab.com:kelda-hotrod/hotrod-base.git
RUN git clone git@gitlab.com:kelda-hotrod/hotrod-driver.git

WORKDIR /go/src/gitlab.com/kelda-hotrod/hotrod-driver

RUN go build -o hotrod main.go
RUN mv hotrod /go/bin/

ENTRYPOINT ["/go/bin/hotrod", "driver"]