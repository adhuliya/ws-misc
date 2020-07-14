# prepare the system
FROM lazynintel/py3-django-uwsgi-nginx:1.0

# the main app location: /app/app

# copy app
WORKDIR /app
COPY . .
RUN chown -R www-data:www-data app/production/runtime

WORKDIR /app/app/production
RUN bash setup_servers.sh

# Start the main process
WORKDIR /app
CMD ["bash", "./start.sh"]
