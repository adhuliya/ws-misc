#!/usr/bin/env bash

# run the docker image -- takes one argument - the tag

NAME="mysite";
ITSOFLIFE="/home/codeman/.itsoflife";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="1.0";
fi

docker run \
  --detach \
  --publish 8080:80 \
  --publish 5055:5050 \
  --mount type=bind,source=$ITSOFLIFE/knots-git,target=/app/knots-git,readonly \
  --mount type=bind,source=$ITSOFLIFE/mydata,target=/app/mydata \
  --name $NAME \
  $NAME:${TAG};


