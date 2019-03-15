FROM golang:1.12.0-stretch

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/marugoshi/gobm

RUN apt-get update -qq \
    && apt-get install -y --no-install-recommends \
    default-libmysqlclient-dev \
    mysql-client \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY Gopkg.toml Gopkg.lock ./

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -v -vendor-only

RUN cd /go/bin \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.2.5/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 migrate \
    && cd /go/src/github.com/marugoshi/gobm

COPY . .