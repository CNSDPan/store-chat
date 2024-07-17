#!/bin/bash
docker exec go-store-chat /bin/bash -c "sh /var/www/store-chat/docker/build/dev-build.sh"
#重启容器
docker restart rpc.socket
docker restart websocket1
docker restart http.api