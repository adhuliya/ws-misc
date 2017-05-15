#!/usr/bin/env python3

import logging as log

from logging.handlers import RotatingFileHandler

"""
Initialize every module with:

    import logging
    log = logging.getLogger(__name__)

Logging levels:

    CRITICAL 50 log.critical()
    ERROR 40 log.error()
    WARNING 30 log.warn()
    INFO 20 log.info()
    DEBUG 10 log.debug()
    NOTSET 0
"""

"""

Edit these three vairables.
"""
logfilename = "main.log"

loglevel = log.INFO # bocks out lower levels

logformat1 = (">>> %(asctime)s : %(levelname)8s : %(filename)s\n"
             "    Line %(lineno)4s : %(funcName)s()\n"
             "%(message)s\n")

logformat2 = (">>> %(asctime)s : %(levelname)8s : %(name)s\n"
             "    Line %(lineno)4s : %(funcName)s()\n"
             "%(message)s\n")

logmaxsize = 1 << 24 # in bytes 1 << 24 = 16 MB
logbackupcount = 5 # 5 x 16MB = 80 MB logs + one extra current 16 MB log.

initialized = False

def init_logger(log_file=logfilename):
    """
    It is inside a function just to see a function name in logger.
    """
    global initialized

    if initialized : return

    #log.basicConfig(format=logformat1)

    rootlogger = log.getLogger()
    rootlogger.setLevel(loglevel)

    handler = RotatingFileHandler(log_file, maxBytes=logmaxsize, backupCount=logbackupcount)
    handler.setFormatter(log.Formatter(logformat2))

    rootlogger.addHandler(handler)

    rootlogger.info("Initialized with format : %s", repr(logformat2))

    initialized = True


if __name__ == "__main__":
    init_logger()


