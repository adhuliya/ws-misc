Development Notes
=======================





## How to deploy using `uWSGI`?

    uwsgi --http :8000 --module mysite.wsgi

This creates/enables the following interaction sequence,

    the web client <-> uWSGI <-> Django
    
## How to add custom error pages?

<https://docs.djangoproject.com/en/3.0/topics/http/views/#customizing-error-views>
