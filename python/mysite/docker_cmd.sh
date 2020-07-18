#!/usr/bin/env bash

# Start all the necessary processes
# Runs inside docker container

# set environment variables
# shellcheck disable=SC2155
export CONTAINER_ID=$(basename "$(cat /proc/1/cpuset)")


# redirect nginx logs to stdout
rm /var/log/nginx/access.log /var/log/nginx/error.log;
ln -s /dev/stdout /var/log/nginx/access.log;
ln -s /dev/stdout /var/log/nginx/error.log;


cd /app/app/production || exit;
bash restart_servers.sh;


# run django db migrations and other django setups
cd /app/app || exit;
python3 manage.py makemigrations;
python3 manage.py migrate;

# Set the admin just once:
# python manage.py createsuperuser --email contact@itsoflife.com --username admin

# Finally wait infinitely. (for testing etc.)
while true; do
  echo "start_sh: `date`";
  sleep 20;
done


