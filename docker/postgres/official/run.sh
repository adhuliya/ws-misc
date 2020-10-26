#!/usr/bin/env bash

# run the docker image -- takes one argument - the tag

NAME="postgres";

if [[ $1 != "" ]]; then
  TAG="$1";
else
  TAG="12.4";
fi

docker run \
  --detach \
  --network="iol" \
  --publish 5432:5432 \
  -e PG_PASSWD_POSTGRES=pg_passwd_postgres.secret \
  -e PG_PASSWD_ADMIN=pg_passwd_admin.secret \
  -e PG_PASSWD_USER1=pg_passwd_user1.secret \
  --name iol.$NAME \
  --mount type=bind,source="$(pwd)/secrets",target=/app/secrets \
  iol.$NAME:${TAG};

sleep 2;

cat ./init_setup.sql | docker exec -i postgres psql -U postgres;

# Unused options:
  # --network="host" \
  #--publish 5432:5432 \


