## ビルドステージ
FROM --platform=linux/amd64 golang:1.21-alpine AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -tags netgo -ldflags '-s -w' -o app
CMD ["./app"]
