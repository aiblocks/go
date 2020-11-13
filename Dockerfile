FROM golang:1.14
WORKDIR /go/src/github.com/aiblocks/go

COPY . .
ENV GO111MODULE=on
RUN go install github.com/aiblocks/go/tools/...
RUN go install github.com/aiblocks/go/services/...
