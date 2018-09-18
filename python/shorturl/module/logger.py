#!/usr/bin/env python3


import subprocess
import os
import os.path as osp
import logging

from logging.handlers import RotatingFileHandler

"""
Usage:
  STEP 1: Only during application initialization,
    import module.logger
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
  """
  Writing logging levels here for convenience of reference.
  """
  CRITICAL = logging.CRITICAL
  ERROR = logging.ERROR
  WARNING = logging.WARNING
  INFO = logging.INFO
  DEBUG = logging.DEBUG
  NOTSET = logging.NOTSET


"""
Edit these configuration vairables:
"""

APP_NAME = None
LOG_DIR = None
LOG_FILE = None

LOG_LEVEL = logging.INFO  # bocks out lower levels

LOG_FORMAT1 = (">>> %(asctime)s : %(levelname)8s : %(filename)s\n"
              "    Line %(lineno)4s : %(funcName)s()\n"
              "%(message)s\n")

LOG_FORMAT2 = (">>> %(asctime)s : %(levelname)8s : %(name)s"
              "    Line %(lineno)4s : %(funcName)s()\n"
              "%(message)s")

LOG_MAX_SIZE = 1 << 24  # in bytes 1 << 24 = 16 MB
LOG_BKP_COUNT = 5  # 5 x 16MB = 80 MB logs + one extra current 16 MB logfile.

INITIALIZED = False


def init_logger(app_name="app", log_file="app.log", log_dir=".", log_level=LOG_LEVEL, log_format=LOG_FORMAT2, log_max_size=LOG_MAX_SIZE, log_bkp_count=LOG_BKP_COUNT):
  global INITIALIZED, LOG_DIR, LOG_FILE, APP_NAME
  if INITIALIZED: return

  APP_NAME = app_name
  LOG_DIR = log_dir
  LOG_FILE = log_file

  log_file = osp.join(log_dir, log_file)
  print("{}: setting up logging system.".format(APP_NAME))

  rootLogger = logging.getLogger()
  rootLogger.setLevel(log_level)

  handler = RotatingFileHandler(log_file, maxBytes=LOG_MAX_SIZE, backupCount=log_bkp_count)
  handler.setFormatter(logging.Formatter(log_format))

  rootLogger.addHandler(handler)

  rootLogger.info("Initialized with format : %s", repr(log_format))

  INITIALIZED = True
  return True

def disable():
  if INITIALIZED:
    logger = logging.getLogger(__name__)
    logger.disabled = True  # and False to enable

def enable():
  if INITIALIZED:
    logger = logging.getLogger(__name__)
    logger.disabled = False  # and True to disable


if __name__ == "__main__":
  init_logger()


