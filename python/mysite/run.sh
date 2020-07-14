#!/usr/bin/env bash

# run the docker image -- takes one argument - the tag

NAME="mysite";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="1.0";
fi

docker run --publish 8080:80 --publish 5055:5050 --detach --name $NAME $NAME:${TAG};
