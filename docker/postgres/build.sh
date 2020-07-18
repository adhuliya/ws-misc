#!/usr/bin/env bash

# build the docker image -- takes one argument - the tag

NAME="postgres";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="13.0";
fi

docker build --tag $NAME:${TAG} .;
