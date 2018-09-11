# a set of utility functions

from typing import Dict, Tuple, List
from .person import Person
from PIL import Image
import os
import os.path as osp
import sys
import io
import re
import time
import string
import random as rand
import bottle as b
import logging
log = logging.getLogger(__name__)

from datetime import datetime
from functools import wraps

# Must initialize these values from a a main code.
# The values here are the defaults, if not overwritten
SESSION_TIMEOUT_SEC = 600  # seconds, must initialize from main code
AVATAR_IMG_SIZE = (256, 256)  # pixels
DEFAULT_IMAGE_QUALITY  = 30  # 30% percent of original

# START default directories
# set them all from the main code
# read here for their relative significance

ROOT_DIR = "/home/codeman/.itsoflife"
TMP_DIR = os.path.join(ROOT_DIR, "local/tmp")
STATIC_DIR = os.path.join(ROOT_DIR, "static")
PEOPLE_DIR = os.path.join(ROOT_DIR, "static/image/db")
IMAGE_DIR = os.path.join(ROOT_DIR, "static/image") 

# path in/for web pages
WEB_PEOPLE_DIR = "db"
WEB_FREEPIC_DIR = "free"
WEB_FILLINS_DIR = "free/fillins"

# END   default directories


def backupdb(people):
    # TODO
    pass

def readdb():
    people = None
    return people

def create_passwd_dict(people : Dict[int, Person]) -> Dict[str, Tuple[int, str]]:
  d = {}
  for k in people:
    person = people[k]
    d[person.email_id] = (k, person.passwd)

  return d

def create_names_list(people : Dict[int, Person]) -> List[Tuple[str, str, int]]:
  user_names = []
  for person in people:
    user_names.append([person.name, person.nick_name, person.id])
  return user_names

def is_correct_login(email_id, passwd, all_users) -> int:
  """checks users's login id password attempt"""
  if email_id in all_users:
    val = all_users[email_id] 
    if val[1] == passwd:
      return val[0]  # the person id
  return None


def is_logged_in(request, online_users):
  """Returns valid userid, if user is logged in, else None"""
  email_id = request.cookies.get("email_id", "")
  if email_id in online_users:
    now = time.time()
    session_info = online_users[email_id]
    activity_gap = now - session_info[1]
    if activity_gap < SESSION_TIMEOUT_SEC:
      online_users[email_id] = (session_info[0], now)
      log.info("good user '%s', id %s", email_id, session_info[0])
      return session_info[0]  # user_id
    else:
      del online_users[email_id]

  return None

def gen_random_string(length=10):
  """returns a random file name, without extension, of the given length"""
  return ''.join(rand.sample(string.ascii_lowercase, length))

def compress_image(filename, quality):
  tmp_name = gen_random_string()
  out_file = tmp_name + ".jpg"
  
  try:
    im = Image.open(filename)
    im.save(out_file, optimize=True, quality=quality)
    return out_file
  except Exception as e:
    log.error("Converting: '%s' TO '%s' failed with: %s", filename, out_file, e)

  return None

def rezise_image(filename, size):
  tmp_name = gen_random_string()
  out_file = tmp_name + ".jpg"
  
  try:
    im = Image.open(filename)
    im.thumbnail(size, Image.ANTIALIAS)
    im.save(out_file, optimize=True, quality=quality)
    return out_file
  except Exception as e:
    log.error("Converting: '%s' TO '%s' failed with: %s", filename, out_file, e)

  return None

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

def fetch_family(person, people):
  family = []
  if person.spouse_id is not None:
    pp = people[person.spouse_id]
    family.append((pp, "spouse"))
  if person.mother_id is not None:
    pp = people[person.mother_id]
    family.append((pp, "mother"))
  if person.father_id is not None:
    pp = people[person.father_id]
    family.append((pp, "father"))

  return family

#    pd.relatives = util.fetch_family(person_id, people)
#    pd.profile_pic = util.profile_pic_path(pd.person)
#    pd.avatar_pic = util.avatar_pic_path(pd.person)
#    pd.pics = util.pics_path(pd.person)

def person_dir(person_id):
  # 100 people data in each directory
  quo = person_id // 100
  low = 100 * quo
  high = 100 * (quo + 1)
  return "{:06}-{:06}".format(low+1, high)

def profile_pic_path(pic_dir, person_id):
  pic_name = "{pid:06}-profile.jpg".format(pid=person_id)
  pth = osp.join(WEB_PEOPLE_DIR, pic_dir, pic_name)
  return pth

def avatar_pic_path(pic_dir, person_id):
  pic_name = "{pid:06}-avatar.jpg".format(pid=person_id)
  pth = osp.join(WEB_PEOPLE_DIR, pic_dir, pic_name)
  return pth

def pics_path(pic_dir, person):
  pics = []
  if person.pics is not None:
    for pic in person.pics:
      pics.append(osp.join(WEB_PEOPLE_DIR, pic_dir, pic))

  return pics

def random_avatar_profile_pic():
  """returns random avatar and profile pics tuple"""
  num = rand.randint(0, 15)
  a = osp.join(WEB_FILLINS_DIR, "{:03}-avatar.jpg".format(num)) 
  p = osp.join(WEB_FILLINS_DIR, "{:03}-profile.jpg".format(num)) 
  return (a,p)

def assign_pics_to_people(people):
  for person in people.values():
    p_dir = person_dir(person.person_id)
    (aa, pp) = random_avatar_profile_pic()

    avatar_pic = avatar_pic_path(p_dir, person.person_id)

    if osp.exists(osp.join(IMAGE_DIR, avatar_pic)):
      person.avatar_pic = avatar_pic
    else:
      log.info("avatar pic = %s, %s", avatar_pic, osp.join(IMAGE_DIR, avatar_pic))
      person.avatar_pic = aa

    profile_pic = profile_pic_path(p_dir, person.person_id)
    if osp.exists(osp.join(IMAGE_DIR, profile_pic)):
      person.profile_pic = profile_pic
    else:
      log.info("profile pic = %s, %s", profile_pic, osp.join(IMAGE_DIR, profile_pic))
      person.profile_pic = pp

  # TODO


def search_name(name, people):
  tmp = [] # list of matches [(score, person)]
  p = build_search_pattern(name)
  for person in people.values():
    finds = p.findall(person.name)
    if finds:
      # magnify mx result by a factor of four
      mx = len(max(finds, key=len)) << 2
      ln = len(finds)
      if mx >= 8  or ln >= 1:
        tmp.append((ln+mx, person))

  tmp.sort(key=lambda x: -x[0])
  result = [[tup[1].person_id, tup[1].name, tup[0]] for tup in tmp[:5]]
  return result

def build_search_pattern(string):
  """build a search pattern from the given string"""
  sio = io.StringIO()

  words = string.lower().split()
  sio.write("|".join(words))
  sio.write("|")

  for word in words:
    if len(word) >= 4:
      for i in range(0, len(word)-3, 4):
        # letter quadruples
        sio.write(word[i:i+4] + "|")
    if len(word) >= 3:
      for i in range(0, len(word)-2, 3):
        # letter triplets
        sio.write(word[i:i+3] + "|")
    for i in range(0, len(word)-1, 2):
      # letter duplets
      sio.write(word[i:i+2] + "|")
  sio.write(";") # any non-matching char

  return re.compile(sio.getvalue(), re.IGNORECASE)

# Only for reference. this is how to return json data.
# def return_json(d):
#     b.response.content_type = "application/json"
#     return dumps(d)

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
        logger.info('%s %s %s %s',b.request.remote_addr,
                                     # request_time,
                                     b.request.method,
                                     b.request.url,
                                     b.response.status)
        return actual_response
    return _log_to_logger

