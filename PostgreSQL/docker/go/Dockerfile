FROM golang:1.19.1-alpine

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go install github.com/cespare/reflex@latest
RUN github.com/kyleconroy/sqlc
ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY go.mod go.sum main.go ./

RUN go mod download
