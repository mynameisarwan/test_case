# syntax=docker/dockerfile:1

FROM golang:latest

RUN mkdir /build
WORKDIR /buidl

RUN export GO111MODULE=on
RUN go get github.com/mynameisarwan/test_case
RUN cd /buid && git clone https://github.com/mynameisarwan/test_case.git

RUN cd /buid/test_case && go build


EXPOSE 8000

ENTRYPOINT [ "/build/MyWeb/test_case/main" ]