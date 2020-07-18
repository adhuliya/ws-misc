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

sleep 2;

cat ./init_setup.sql | docker exec -i postgres psql -U postgres;

# Unused options:
  # --network="host" \
  #--publish 5432:5432 \


