Development Notes
=======================





## How to deploy using `uWSGI`?

    uwsgi --http :8000 --module mysite.wsgi

This creates/enables the following interaction sequence,

    the web client <-> uWSGI <-> Django
