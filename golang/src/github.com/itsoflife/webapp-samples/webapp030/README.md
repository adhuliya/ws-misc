Basic helloworld App
=========================

This is a simple hello world webapp with a properly configured seelog logger, and simple examples on how to use it.

The initialization code for the logger is stored in `./vendor/app/logs/logconfig.go`. `vendor` is a special folder used to put all application code. Always put all the project packages in `./vendor/app` and see the `./main.go` file on how to include these packages by simply writing `import "app/logs"`. The log output is stored in `./logs` folder.

How to run?
-----------------

    go run main.go


Dependencies
======================

* seelog

        go get -u github.com/cihub/seelog

  Its configuration is in ./configs/seelog.xml file.
