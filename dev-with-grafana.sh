#!/bin/bash

set -e

COMPOSE="docker compose"

if [[ -x "$(command -v docker-compose)" ]]; then
    COMPOSE=docker-compose
fi

$COMPOSE --profile grafana up --build $@
