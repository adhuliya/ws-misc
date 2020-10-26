#!/usr/bin/env bash

# REF: https://www.postgresql.org/download/linux/ubuntu/

# Create the file repository configuration:
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list';

# Import the repository signing key:
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -;

# Update the package lists:
sudo apt-get update;

# Install the latest version of PostgreSQL.
# If you want a specific version, use 'postgresql-12' or similar instead of 'postgresql':
sudo apt-get -y install postgresql-12;
sudo apt-get -y install postgresql-client-12;
sudo apt-get -y install pgadmin4;

# List of packages available:
#   postgresql-client-12 	  client libraries and client binaries
#   postgresql-12 	        core database server
#   postgresql-contrib-9.x 	additional supplied modules (part of the postgresql-xx package in version 10 and later)
#   libpq-dev 	            libraries and headers for C language frontend development
#   postgresql-server-dev-12 	libraries and headers for C language backend development
#   pgadmin4                pgAdmin 4 graphical administration utility
