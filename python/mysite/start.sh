#!/usr/bin/env bash

# Start all the necessary processes
# Runs inside docker container

cd /app/app/production;
bash restart_servers.sh;

# Finally wait infinitely. (for testing etc.)
while true; do
  echo "Hello `date`";
  sleep 10;
done
