FROM golang:1.19-alpine as builder

# Default workspace.
RUN mkdir -p /go/src/workspace/perso/custom-fizzbuzz
COPY . /go/src/workspace/perso/custom-fizzbuzz

ARG GIT_TAG_NAME=wip
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=mod
ENV GOOS=linux
ENV GOARCH=amd64


WORKDIR /go/src/workspace/perso/custom-fizzbuzz
RUN go mod download

RUN go build -ldflags "-X main.buildVersion=v0.0.0-docker" -o /app ./cmd/

WORKDIR /
CMD  ./app