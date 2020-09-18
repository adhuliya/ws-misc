Development Notes
=======================

This is a mother project which integrates many sub-projects
that demonstrate and use important features that are useful
in web development.


## TODO

1. Run and test the project - json and protected file etc. - DONE
2. Add login capability as an app.
3. Add app to display all user agent data. -- DONE
4. Add pretty/informative error pages.
5. Add logging capability -- DONE
6. Add REST API and learn to test it with curl.
7. Add template inheritance. -- DONE basic.
8. Add Twilio SMS capability.
9. Add push notification capability.
10. Add email sending/receiving capability.
11. Add docker image for the django server system.
12. What is the advantage of Postgres in a docker?

## How to deploy using `uWSGI`?

    uwsgi --http :8000 --module mysite.wsgi

This creates/enables the following interaction sequence,

    the web client <-> uWSGI <-> Django
    
## How to add custom error pages?

<https://docs.djangoproject.com/en/3.0/topics/http/views/#customizing-error-views>


