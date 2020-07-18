#!/usr/bin/env bash

# remove the docker container -- takes one argument - the container name

NAME="postgres";

if [[ $1 != "" ]]; then
  NAME="$1";
fi

docker rm --force $NAME;
