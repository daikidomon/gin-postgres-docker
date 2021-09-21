FROM golang:1.17-buster

RUN mkdir /go/src/app
RUN mkdir -p /data
RUN apt-get update \
  && apt-get install git

RUN go get -u -t github.com/gin-gonic/gin

WORKDIR /go/src/app