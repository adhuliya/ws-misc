#!/usr/bin/env bash

# Needs sudo permission.
# This scripts restarts nginx and uwsgi servers.
# setup_servers.sh does the initial setup of the servers.

# STEP 1: restart uwsgi servers

echo "Restarting uwsgi...";
if [[ -e runtime/uwsgi.pid ]]; then
  /usr/local/bin/uwsgi --stop runtime/uwsgi.pid;
  #/usr/local/bin/uwsgi --reload runtime/uwsgi.pid;
else
  echo "runtime/uwsgi.pid file not found. Going for a brutal `killall uwsgi`!";
  killall uwsgi;
  killall uwsgi;  # it needs two kills in quick successions !!
fi
sleep 2; # wait to make sure they die
/usr/local/bin/uwsgi --emperor /etc/uwsgi/vassals \
  --uid www-data --gid www-data --daemonize /var/log/uwsgi-emperor.log;

# STEP 2: restart nginx server

echo "Restarting nginx...";
systemctl restart nginx.service;

echo "Done.";

