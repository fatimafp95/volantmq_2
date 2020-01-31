#!/usr/bin/env bash

buildArgs=

if [[ ! -z "$1" ]]; then
    buildArgs="-t $1"
fi

docker build -f ./local.Dockerfile -t volantmq2local ${buildArgs} .

echo "removing intermediate container"
docker rmi -f $(docker images -q --filter label=stage=intermediate)
