#!/bin/bash

set -e
cd ${0%/*}
PREFIX=ghcr.io/kangaroux

docker build -t $PREFIX/booruview-api -f api/Dockerfile.prod api/
docker build -t $PREFIX/booruview-valkey valkey/
docker build -t $PREFIX/booruview-client client/

OWNER=$(id -u):$(id -g)
docker run --rm -e OWNER=$OWNER -v ./client/dist:/app/dist booruview-client ash -c 'yarn build && chown -R $OWNER dist/'
docker build -t $PREFIX/booruview-caddy -f caddy/Dockerfile .

docker push $PREFIX/booruview-api
docker push $PREFIX/booruview-caddy
docker push $PREFIX/booruview-valkey
