#!/usr/bin/env python3
"""
Global settings for the `shorten_url` project.
Importing this module executes the settings.
"""

# no space in app name
APP_NAME = "shorten_url"

import os
import os.path as osp
import logging

from . import logger
from .util import create_dir,\
        file_exists

# relative paths use user's home directory
APP_DIR = ".itsoflife/local/.apps/" + APP_NAME

DATA_DIR = osp.join(APP_DIR, "data")
DATA_USER_FILE = "user_map.p"
DATA_URL_FILE = "url_map.p"

LOG_FILE = APP_NAME + ".log"
LOG_DIR = osp.join(APP_DIR, "logs")


# the working directory of the program
PROG_DIR = "..."


################################################
# enforce settings
################################################

LOG_DIR = file_exists(LOG_DIR, is_dir=True) or create_dir(LOG_DIR)
assert LOG_DIR, "Failed creating Logging Dir: {}".format(LOG_DIR)
logger.init_logger(app_name=APP_NAME, log_file=LOG_FILE, log_dir=LOG_DIR)

log = logging.getLogger(__name__)

APP_DIR = file_exists(APP_DIR, is_dir=True) or create_dir(APP_DIR)
assert APP_DIR

DATA_DIR = file_exists(DATA_DIR, is_dir=True) or create_dir(DATA_DIR)
assert DATA_DIR

################################################

