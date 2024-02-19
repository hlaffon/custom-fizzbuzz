FROM golang:1.21

RUN mkdir -p ${HOME}/custom-fizzbuzz
COPY . ${HOME}/custom-fizzbuzz

ENV CGO_ENABLED 1
ENV GOFLAGS -mod=mod
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR ${HOME}/custom-fizzbuzz
RUN go mod download

RUN go build -ldflags "-X main.buildVersion=v0.0.0-docker" -o /app ./cmd/

WORKDIR /
CMD  ./app