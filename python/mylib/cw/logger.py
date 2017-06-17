#!/usr/bin/env python3

import logging

from logging.handlers import RotatingFileHandler

"""
Usage:
  STEP 1: Only during application initialization,
    import cw.logger
    logger.init_logger()

  STEP 2: For each module in the application,
    import logging
    log = logging.getLogger(__name__)

Logging levels:

    CRITICAL    50 log.critical()
    ERROR       40 log.error()
    WARNING     30 log.warn()
    INFO        20 log.info()
    DEBUG       10 log.debug()
    NOTSET      0
"""

class LogLevels:
    CRITICAL    = logging.CRITICAL
    ERROR       = logging.ERROR
    WARNING     = logging.WARNING
    INFO        = logging.INFO
    DEBUG       = logging.DEBUG
    NOTSET      = logging.NOTSET


"""
Edit these configuration vairables:
"""
logFileName = "main.log"

logLevel = logging.INFO # bocks out lower levels

logFormat1 = (">>> %(asctime)s : %(levelname)8s : %(filename)s\n"
             "    Line %(lineno)4s : %(funcName)s()\n"
             "%(message)s\n")

logFormat2 = (">>> %(asctime)s : %(levelname)8s : %(name)s\n"
             "    Line %(lineno)4s : %(funcName)s()\n"
             "%(message)s\n")

logMaxSize = 1 << 24 # in bytes 1 << 24 = 16 MB
logBackupCount = 5 # 5 x 16MB = 80 MB logs + one extra current 16 MB logfile.

initialized = False

def initLogger(logFile=logFileName, logLevel=logLevel, logFormat=logFormat2, logMaxSize=logMaxSize, logBackupCount=logBackupCount):

    global initialized
    if initialized : return

    #logging.basicConfig(format=logformat1)

    rootLogger = logging.getLogger()
    rootLogger.setLevel(logLevel)

    handler = RotatingFileHandler(logFile, maxBytes=logMaxSize, backupCount=logBackupCount)
    handler.setFormatter(logging.Formatter(logFormat))

    rootLogger.addHandler(handler)

    rootLogger.info("Initialized with format : %s", repr(logFormat))

    initialized = True


if __name__ == "__main__":
    initLogger()


