#!/usr/bin/env bash

# sudo permission needed.
# Setup the nginx and uwsgi servers.
# The working directory of this script should contain the following files,
#   * mysite_nginx.conf
#   * mysite_uwsgi.ini
#   * uwsgi_params  -- referred in mysite_nginx.conf file

# NOTE: absoulute path in both these files needs to be manually updated.
NGINX_CONF=mysite_nginx.conf;
UWSGI_CONF=mysite_uwsgi.ini;

PWD="`pwd`";

NGINX_CONF_DIR="/etc/nginx/sites-enabled";
UWSGI_CONF_DIR="/etc/uwsgi/vassals";

# Setup nginx configuration
NGINX_DEST_FILE=$NGINX_CONF_DIR/$NGINX_CONF;
if [[ -h $NGINX_DEST_FILE ]]; then
  echo "Removing the old $NGINX_DEST_FILE";
  rm $NGINX_DEST_FILE;  # remove the old symlink
else
  if [[ -e $NGINX_DEST_FILE ]]; then
    echo "ERROR: Exists and not a symlink: $NGINX_DEST_FILE";
    exit 1;
  fi
fi

ln -s "$PWD/$NGINX_CONF" "$NGINX_DEST_FILE";

# Setup uwsgi configuration
mkdir -p "$UWSGI_CONF_DIR";

UWSGI_DEST_FILE=$UWSGI_CONF_DIR/$UWSGI_CONF;
if [[ -h $UWSGI_DEST_FILE ]]; then
  echo "Removing the old $UWSGI_DEST_FILE";
  rm $UWSGI_DEST_FILE;  # remove the old symlink
else
  if [[ -e $UWSGI_DEST_FILE ]]; then
    echo "ERROR: Exists and not a symlink: $UWSGI_DEST_FILE";
    exit 2;
  fi
fi

ln -s "$PWD/$UWSGI_CONF" "$UWSGI_DEST_FILE";

# To start the servers see: restart_servers.sh



