# prepare the system
FROM python:3.6-buster

# on the python image add nginx and uwsgi
RUN apt-get update
RUN apt-get install -y nginx
RUN pip3 install uwsgi
RUN pip3 install django

# adding bottle for mini tests
RUN pip3 install bottle

# install postgres driver for python
RUN pip3 install psycopg2

#DONE

