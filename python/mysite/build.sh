#!/usr/bin/env bash

# build the docker image -- takes one argument - the tag

NAME="mysite";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="1.0";
fi

docker build --tag $NAME:${TAG} .;
