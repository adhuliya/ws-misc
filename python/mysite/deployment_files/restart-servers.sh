#!/usr/bin/env bash

# Needs sudo permission.
# This scripts restarts nginx and uwsgi servers.

# STEP 1: restart uwsgi servers

echo "Restarting uwsgi...";
killall uwsgi;
sleep 4;
/usr/local/bin/uwsgi --emperor /etc/uwsgi/vassals --uid www-data --gid www-data --daemonize /var/log/uwsgi-emperor.log;

# STEP 2: restart nginx server

echo "Restarting nginx...";
systemctl restart nginx.service;

echo "Done.";

