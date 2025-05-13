#!/bin/bash

set -e
cd ${0%/*}
PREFIX=ghcr.io/kangaroux
COMMIT=$(git rev-parse --short master)
DATE=$(date +%Y.%m.%d)
TAG=$COMMIT-$DATE

docker build -t $PREFIX/booruview-api:$TAG -f api/Dockerfile.prod api/
docker build -t $PREFIX/booruview-valkey:$TAG valkey/
docker build -t $PREFIX/booruview-client client/

OWNER=$(id -u):$(id -g)
docker run --rm -e OWNER=$OWNER -v ./client/dist:/app/dist booruview-client ash -c 'yarn build && chown -R $OWNER dist/'
docker build -t $PREFIX/booruview-caddy:$TAG -f caddy/Dockerfile .

docker push $PREFIX/booruview-api:$TAG
docker push $PREFIX/booruview-caddy:$TAG
docker push $PREFIX/booruview-valkey:$TAG
