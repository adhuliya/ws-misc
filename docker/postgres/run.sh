#!/usr/bin/env bash

# run the docker image -- takes one argument - the tag

NAME="postgres";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="13.0";
fi

docker run \
  --detach \
  --network="itsoflife" \
  --publish 5432:5432 \
  -e POSTGRES_PASSWORD=anshuisneo \
  --name $NAME \
  $NAME:${TAG};
# Unused options:
  # --network="host" \
  #--publish 5432:5432 \


