URL Shortner and Forwarder
===============

Forward short urls!

TODO
-------
1. add a form

How to start?
-----------------

    python3 main.py

How to speed test?
-----------------
Record the start and end time for the following:

    for x in {1..10000}; do wget -q -O /dev/null http://localhost:5000/test ; done &> /dev/null 

Implementation Notes
----------------------
* `./views` is the default templates directory
