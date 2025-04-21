#!/bin/sh

cd ${0%/*}

cat << EOF > .env
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=$DBURL
GOOSE_MIGRATION_DIR=./migrations
EOF
