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

NGINX_CONF_DIR="/etc/nginx/sites-enabled";
UWSGI_CONF_DIR="/etc/uwsgi/vassals";

PWD="`pwd`";

function prepareDestination () {
  # Takes one argument: The absolute destination file path.
  # This function removes the old symlink if it exists,
  # or exits with an error if a file with same name exits
  # and is not symlink as exptected.
  local DEST_FILE=$1;
  if [[ -h $DEST_FILE ]]; then
    echo "Removing the old $DEST_FILE";
    rm $DEST_FILE;  # remove the old symlink
  else
    if [[ -e $DEST_FILE ]]; then
      echo "ERROR: Exists and not a symlink: $DEST_FILE";
      exit 1; # exits this script
    fi
  fi
}

# Setup nginx configuration
NGINX_DEST_FILE=$NGINX_CONF_DIR/$NGINX_CONF;
prepareDestination "$NGINX_DEST_FILE";
ln -s "$PWD/$NGINX_CONF" "$NGINX_DEST_FILE";

# Setup uwsgi configuration
mkdir -p "$UWSGI_CONF_DIR";
UWSGI_DEST_FILE=$UWSGI_CONF_DIR/$UWSGI_CONF;
prepareDestination "$UWSGI_DEST_FILE";
ln -s "$PWD/$UWSGI_CONF" "$UWSGI_DEST_FILE";

# To start the servers see: restart_servers.sh



