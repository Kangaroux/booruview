#!/bin/bash

set -e
cd ${0%/*}

echo "Building for production"

docker build -t booruview-api -f api/Dockerfile.prod api/
docker build -t booruview-valkey valkey/
docker build -t booruview-client client/

OWNER=$(id -u):$(id -g)
docker run --rm -e OWNER=$OWNER -v ./client/dist:/app/dist booruview-client ash -c 'yarn build && chown -R $OWNER dist/'
docker build -t booruview-caddy -f caddy/Dockerfile .

echo "Done."
