#!/bin/bash
cd /var/www/store-chat/socket-client/ &&
# 设置打包环境
GOOS=linux GOARCH=amd64 go build -o bin/socket.client.bin -tags=socket.client client.go