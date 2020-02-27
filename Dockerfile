# ARG給參數預設值 如果``` $docker build --build-arg GO_VERSION=1.13 ```未指定新的值，就使用先指定好的值
ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./src/main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .
RUN mkdir -p config
COPY --from=builder /api/src/config/ ./config/
#COPY --from=builder /api/test.db .

EXPOSE 8080

ENTRYPOINT ["./app"]