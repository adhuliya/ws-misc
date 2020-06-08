Development Notes
=======================



## TODO

1. Run and test the project - json and protected file etc.
2. Add login capability as an app. -- and test protected file feature
3. Add app to print all user agent data.
4. Add informative error pages.
5. Add logging capability
6. Add REST API and learn to test it with curl.
7. Add template inheritance.


## How to deploy using `uWSGI`?

    uwsgi --http :8000 --module mysite.wsgi

This creates/enables the following interaction sequence,

    the web client <-> uWSGI <-> Django
    
## How to add custom error pages?

<https://docs.djangoproject.com/en/3.0/topics/http/views/#customizing-error-views>
