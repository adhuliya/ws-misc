#!/usr/bin/env bash

NAME="mysite";

docker exec -w /app/app -it $NAME  python3 manage.py createsuperuser --email contact@itsoflife.com --username admin 
