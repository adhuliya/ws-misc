#!/usr/bin/env bash

# run the docker image -- takes one argument - the tag
# Names starting with `_` are not exported.

_NAME="mysite";
_ITSOFLIFE="/home/codeman/.itsoflife";

echo "Removing any container with name $_NAME" >> $MYDATA/local/logs/mysystem.log;
/usr/bin/docker rm --force $_NAME;

if [[ $1 != "" ]]; then
  _TAG="$1";
else
  _TAG="1.0";
fi

echo "Starting container with name $_NAME" >> $MYDATA/local/logs/mysystem.log;
# make sure network is created for --network="itsoflife":
#   docker network create itsoflife;
/usr/bin/docker run \
  --detach \
  --publish 8080:80 \
  --publish 5055:5050 \
  --network="itsoflife" \
  --mount type=bind,source=$_ITSOFLIFE/knots-git,target=/app/knots-git,readonly \
  --mount type=bind,source=$_ITSOFLIFE/mydata,target=/app/mydata \
  --name $_NAME \
  $_NAME:${_TAG};

