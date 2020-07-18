# prepare the system
FROM lazynintel/py3-django-uwsgi-nginx:1.0

# For django rest framework: https://www.django-rest-framework.org
RUN pip3 install djangorestframework
# Markdown support for the browsable API.
RUN pip3 install markdown
# Filtering support
RUN pip3 install django-filter

# the main app location: /app/app

# copy app
WORKDIR /app
COPY . .
RUN chown -R www-data:www-data app/production/runtime

WORKDIR /app/app/production
RUN bash setup_servers.sh

# Start the main process
WORKDIR /app
CMD ["bash", "./docker_cmd.sh"]
