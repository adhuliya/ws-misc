"""
The data containing user details and the urls.
"""
import pickle
import os.path as osp
import logging
log = logging.getLogger(__name__)

from .util import create_dir,\
    file_exists

from .settings import DATA_DIR, DATA_URL_FILE, DATA_USER_FILE

class User():
    def __init__(self,
            user_id = None,
            name = "John Doe",
            email_id = "john@gmail.com",
            passwd = None,
            urls = None):
        self.user_id = user_id
        self.name = name.lower()
        self.email_id = email_id
        self.passwd = passwd
        self.urls = [] if not urls else urls

    def __repr__(self):
        return ("User(user_id={0.user_id}, name={0.name!r},"
                " email_id={0.email_id!r}, passwd={0.passwd!r},"
                " urls={0.urls!r})").format(self)

admin = User(user_id=1, name="lazy nintel",
        email_id="lazynintel@gmail.com", passwd="anshuisneo",
        urls=["asdfghsd", "asdfwewe", "asdfsdfc"])

dummy_urls = {"asdfghsd": "https://www.google.com",
        "asdfwewe": "https://www.yahoo.com",
        "asdfsdfc": "https://www.cse.iitb.ac.in/~dhuliya"}

# map of email_id -> User Object
user_map = dict()

# map of short url code -> full url
url_map = dict()

# empty when server starts
online_users = dict()

class PageData():
    """
    Object used to pass data to templates.
    Use: `<PageData_obj>.__dict__` to pass data
    """
    def __init__(self,
            title = "Shorten URL",
            user = None,
            url_map = None):
        self.title = title
        self.user = user
        self.url_map = url_map


def load_users(user_file_path):
  user_map = None
  with open(user_file_path, "rb") as f:
    user_map = pickle.load(f)
  return user_map

def save_users(user_file_path, user_map):
  with open(user_file_path, "wb") as f:
    pickle.dump(f, user_map)
  return True

def load_urls(url_file_path):
  url_map = None
  with open(url_file_path, "rb") as f:
    url_map = pickle.load(f)
  return url_map

def save_urls(url_file_path, url_map):
  with open(url_file_path, "wb") as f:
    pickle.dump(f, url_map)
  return True

def load_data():
    global user_map, url_map, admin, dummy_urls

    users_file_path = osp.join(DATA_DIR, DATA_USER_FILE)
    if file_exists(users_file_path):
        user_map = load_users(users_file_path)

    if admin.email_id not in user_map:
        user_map[admin.email_id] = admin

    url_file_path = osp.join(DATA_DIR, DATA_URL_FILE)
    if file_exists(url_file_path):
        url_map = load_users(url_file_path)

    for key, val in dummy_urls.items():
        if key not in url_map:
            url_map[key] = val



