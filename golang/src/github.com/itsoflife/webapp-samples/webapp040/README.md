Basic helloworld App
=========================

This is a simple hello world webapp with a properly configured seelog logger, and hierarchical html template system.

The initialization code for the logger is stored in `./vendor/app/logs/logconfig.go`. `vendor` is a special folder used to put all application code. Always put all the project packages in `./vendor/app` and see the `./main.go` file on how to include these packages by simply writing `import "app/logs"`. The log output is stored in `./logs` folder.

The initialization code for the template engine is stored in `./vensor/app/tmpl/template.go`. It expects a particular folder structure in the `./template` directory.

How to run?
-----------------

    go run main.go


Dependencies
======================

* seelog

        go get -u github.com/cihub/seelog

  Its configuration is in ./configs/seelog.xml file.

* bpool

        go get github.com/oxtoacart/bpool

  It used by the template rendering function to first store the rendered page into buffer, so that if an error occurs, the user can be redirected to some other page.


