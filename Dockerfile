# syntax=docker/dockerfile:1

FROM golang:latest

RUN mkdir /build
WORKDIR /buidl

RUN export GO111MODULE=on
RUN go get github.com/mynameisarwan/MyWeb/main

EXPOSE 8000

ENTRYPOINT [ "/build/MyWeb/main/main" ]