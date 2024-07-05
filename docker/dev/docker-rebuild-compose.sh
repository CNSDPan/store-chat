#!/bin/bash
docker exec go-store-chat /bin/bash -c "sh /var/www/store-chat/docker/build/dev-build.sh"
docker-compose -f docker-compose.yml up -d --build