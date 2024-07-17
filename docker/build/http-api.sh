#!/bin/bash
cd /var/www/store-chat/api/ &&
# 设置打包环境
GOOS=linux GOARCH=amd64 go build -o bin/api.bin -tags=api api.go