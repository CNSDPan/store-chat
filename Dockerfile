FROM golang:1.21 AS build
ENV TZ=Asia/Shanghai
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /var/www/store-chat/

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download && go mod verify

RUN apt update&&apt upgrade -y
RUN apt install vim -y