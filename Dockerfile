FROM golang:1.12.0-stretch

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/marugoshi/gobm

COPY Gopkg.toml Gopkg.lock ./

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -v -vendor-only

COPY . .