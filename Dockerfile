FROM golang:1.21.0-alpine3.18 as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy
RUN go build .

FROM docker:dind

RUN apk update && apk upgrade
RUN apk add --no-cache git make musl-dev tmux bash go

RUN adduser -h /home/tmuxer -s /bin/bash -G root -D tmuxer
USER tmuxer

RUN mkdir -p /home/tmuxer/gokube
WORKDIR /home/tmuxer/gokube

COPY --from=builder /build/gokube .
COPY . .

USER root

EXPOSE 22/tcp