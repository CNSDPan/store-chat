#!/bin/bash
docker-compose -f docker-compose-build.yml up -d --build
docker exec go-store-chat-build /bin/bash -c "sh /var/www/store-chat/docker/build/dev-build.sh"
docker rm -f go-store-chat-build
docker-compose -f docker-compose.yml up -d --build