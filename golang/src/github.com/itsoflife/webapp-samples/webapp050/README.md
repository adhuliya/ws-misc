Basic helloworld App
=========================

This is a simple webapp that shows the full working with a postgres database. Other included features are,

1. A properly configured seelog logger.
2. A hierarchical html template system.
3. Automation script to setup database with ease.

How to run?
-----------------

    go run main.go


About Application
-----------------------

This is a simple hello world webapp with a properly configured seelog logger, and hierarchical html template system.

The initialization code for the logger is stored in `./vendor/app/logs/logconfig.go`. `vendor` is a special folder used to put all application code. Always put all the project packages in `./vendor/app` and see the `./main.go` file on how to include these packages by simply writing `import "app/logs"`. The log output is stored in `./logs` folder.

The initialization code for the template engine is stored in `./vendor/app/tmpl/template.go`. It expects a particular folder structure in the `./template` directory.

FAQs
----------------------
### Q. How to setup the database?
Install postgres database and make sure you have the default `postgres` account and database (successful installation ensures that). Then run the following commands,

    $ ## Assuming current directory same as this README.md file.
    $ cd configs/db/postgres
    $ ./setup.py roles
    $ ./setup.py setup

These commands run the sql scripts in the folders inside `configs/db/postgres` automatically in the right order.

Dependencies
======================

* seelog

        go get -u github.com/cihub/seelog

  Its configuration is in ./configs/seelog.xml file.

* bpool

        go get github.com/oxtoacart/bpool

  It is used by the template rendering function to first store the rendered page into buffer, so that if an error occurs, the user can be redirected to some other page.


