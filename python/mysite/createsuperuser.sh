#!/usr/bin/env bash

_NAME="mysite";

docker exec -w /app/app -it $_NAME  python3 manage.py createsuperuser --email contact@itsoflife.com --username admin 
