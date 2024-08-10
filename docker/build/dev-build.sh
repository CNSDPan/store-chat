#!/bin/bash
cd /var/www/store-chat/docker/build &&
./rpc-socket.sh
./socket.sh
./socket-client.sh
./http-api.sh