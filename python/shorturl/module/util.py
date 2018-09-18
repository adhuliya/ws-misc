# a set of utility functions
"""
A collection of STATELESS utility functions.
"""

from typing import Dict, Tuple, List
import os
import os.path as osp
import sys
import io
import re
import time
import string
import random as rand
from json import dumps
from bottle import request, response
import logging
log = logging.getLogger(__name__)

from datetime import datetime
from functools import wraps

letters = string.ascii_letters + string.digits

def delete_short_url(user, url_map, short_url):
  if short_url in user.urls:
    user.urls.remove(short_url)
    if short_url in url_map:
      del url_map[short_url]
      return True
  return False

def is_correct_login(email_id, passwd, all_users) -> str:
  """checks users's login id password attempt"""
  if email_id in all_users:
    user = all_users[email_id] 
    if user.passwd == passwd:
      return email_id
  return None

datacount = 1

def is_logged_in(request, online_users, timeout_sec=600):
  """Returns valid email_id, if user is logged in, else None"""
  global datacount
  datacount += 1
  email_id = request.cookies.get("email_id", "")
  if email_id in online_users:
    now = time.time()
    last_active = online_users[email_id]
    activity_gap = now - last_active
    if activity_gap < timeout_sec:
      online_users[email_id] = now
      log.info("%s, good user '%s', id %s", datacount, email_id, last_active)
      return email_id
    else:
      del online_users[email_id]

  log.info("why here: %s, %s, %s", datacount, email_id, online_users)
  return None

def gen_random_string(length=8):
  """returns a random string, of the given length"""
  return ''.join(rand.sample(letters, length))

# generates and adds a unique short url
def gen_unique_short_url(url_map):
  while True:
    short_url = gen_random_string()
    if short_url in url_map:
      continue
    else:
      break
    # reserver a slot
    url_map[short_url] = None
  return short_url


def remove_file(filename):
  """removes only normal file and returns boolean success"""
  exists = os.path.exists(filename)
  if exists and os.path.isfile(filename):
    try:
      os.remove(filename)
      return True
    except Exception as e:
      log.error("Removing: '%s' failed, error: %s", filename, e)

  return False


# Only for reference. this is how to return json data.
def return_json(d):
  response.content_type = "application/json"
  return dumps(d)

# function to add logging to bottle.py
# Usage:
#   app = Bottle()
#   app.install(log_to_logger)
#   app.run(host='0.0.0.0', port='8080', quiet=True)
# `quiet=True` suppresses output to the console.
def log_to_logger(fn):
    '''
    Wrap a Bottle request so that a log line is emitted after it's handled.
    (This decorator can be extended to take the desired logger as a param.)
    '''
    logger = logging.getLogger(__name__)

    @wraps(fn)
    def _log_to_logger(*args, **kwargs):
        # request_time = datetime.now()
        actual_response = fn(*args, **kwargs)
        # modify this to log exactly what you need:
        logger.info('%s %s %s %s', request.remote_addr,
                                     # request_time,
                                     request.method,
                                     request.url,
                                     response.status)
        return actual_response
    return _log_to_logger


# checks if a file/dir exists, and returns absolute path
# if path is relative its checked in user's home directory
# in case of errors returns None and logs the error.
def file_exists(filepath, is_dir=False):
  abs_path = None
  if osp.isabs(filepath):
    abs_path = filepath
  else:
    user_home = os.getenv("HOME", None).strip()
    if user_home:
      abs_path = osp.join(user_home, filepath)
    else:
      log.error("Unable to check '%s'. Env variable 'HOME' empty /not available.", filepath)
      return None

  exists = False
  
  try:
    exists = osp.exists(abs_path)
    if is_dir:
      exists = osp.isdir(abs_path)
  except Exception as e:
    log.error("Error checking path %s,\n%s", abs_path, e)
    exists = False

  return abs_path if exists else None


# creates the given directory, and returns absolute path
# if path is relative its created in user's home directory
# in case of errors returns None and logs the error.
def create_dir(dirpath):
  abs_path = None
  if osp.isabs(dirpath):
    abs_path = dirpath
  else:
    user_home = os.getenv("HOME", None).strip()
    if user_home:
      abs_path = osp.join(user_home, dirpath)
    else:
      log.error("Unable to create dir '%s'. Env varialbe 'HOME' empty /not available.", dirpath)
      return None
  
  try:
    os.makedirs(abs_path)
  except Exception as e:
    log.error("Error creating directory %s,\n%s", abs_path, e)
    return None

  return abs_path

