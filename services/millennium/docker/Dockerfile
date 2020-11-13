FROM golang:1.14-stretch AS builder

WORKDIR /go/src/github.com/aiblocks/go
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install github.com/aiblocks/go/services/millennium

FROM ubuntu:18.04

ENV AIBLOCKS_CORE_VERSION 14.1.1-1355-5b2989d3
ENV AIBLOCKS_CORE_BINARY_PATH /usr/local/bin/aiblocks-core

# ca-certificates are required to make tls connections
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates wget


RUN apt-get install -y --no-install-recommends libpqxx-4.0v5 curl
RUN wget -O aiblocks-core.deb https://s3.amazonaws.com/aiblocks.io/releases/aiblocks-core/aiblocks-core-${AIBLOCKS_CORE_VERSION}_amd64.deb
RUN dpkg -i aiblocks-core.deb
RUN rm aiblocks-core.deb

RUN apt-get clean

COPY --from=builder /go/bin/millennium ./

ENTRYPOINT ["./millennium"]
