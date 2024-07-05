#!/bin/bash
cd /var/www/store-chat/socket/ &&
# 设置打包环境
GOOS=linux GOARCH=amd64 go build -o bin/socket.bin -tags=socket socket.go