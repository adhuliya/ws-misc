#!/usr/bin/env bash

# build the docker image -- takes one (optional) argument - the tag

NAME="postgres";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="12.4";
fi

docker build --tag iol.$NAME:${TAG} .;
