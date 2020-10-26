#!/usr/bin/env bash

# run the docker image -- takes one (optional) argument - the tag
# Names starting with `_` are not exported.

# For production either don't define DEV_SERVER env variable or
# set it to an empty string. (by default production is assumed)

# MUST run this script as `./run.sh`

if [[ ! -d ../app || ! -d ../production ]]; then
  # a simple check to ensure the invocation is correct.
  echo "AD: Must run this script as `./run.sh`";
  exit;
fi

_HOST_USER=$USER
_HOST_APP="`pwd`/../app";  # the path in host
_APP_NAME="py3django";
_PORT=80

# production server inits
_HOST_PROD=$_HOST_APP/../production;
_HOST_SECRETS=$_HOST_APP/../secrets;
_HOST_KNOTS=/home/admin/knots-git;
_HOST_MYDATA=/home/admin/mydata;

if [[ -n "$DEV_SERVER" ]];
then
  # its a development server
  echo "AD: Going to run DEVELOPMENT server!";
  _HOST_KNOTS=/home/codeman/.itsoflife/knots-git;
  _HOST_MYDATA=/home/codeman/.itsoflife/mydata;
  _PORT=5055
fi

_DOCKER_APP=/app;
_DOCKER_PROD=$_DOCKER_APP/production;
_DOCKER_DATA=$_DOCKER_PROD/data;
_DOCKER_LOGS=$_DOCKER_DATA/logs;
_DOCKER_KNOTS=$_DOCKER_APP/knots-git;
_DOCKER_MYDATA=$_DOCKER_APP/mydata;

# set permissions on the data folder
chown -R www-data:www-data $_HOST_PROD/data;

echo "AD: Removing any container with name $_APP_NAME";
/usr/bin/docker rm --force $_APP_NAME;

if [[ $1 != "" ]]; then
  _TAG="$1";
else
  _TAG="1.0";
fi

echo "AD: Starting container with name $_APP_NAME";
# make sure network is created for --network="itsoflife": (if using it)
#   docker network create itsoflife;
/usr/bin/docker run \
  --detach \
  --publish $_PORT:5050 \
  --mount type=bind,source=$_HOST_APP,target=/app/app,readonly \
  --mount type=bind,source=$_HOST_SECRETS,target=/app/secrets,readonly \
  --mount type=bind,source=$_HOST_PROD,target=/app/production \
  --mount type=bind,source=$_HOST_KNOTS,target=/app/knots-git,readonly \
  --mount type=bind,source=$_HOST_MYDATA,target=/app/mydata,readonly \
  --env DEV_SERVER="$DEV_SERVER" \
  --name $_APP_NAME \
  $_APP_NAME:${_TAG};
  # --network="itsoflife" \
  # --publish 5055:5050 \

