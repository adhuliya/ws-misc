# A simple bottle based web app for family tree
import sys
import os
import os.path as osp
import json

import bottle as b

import module.util as u
from module.person import Person
from module.people import people
from module.page_data import PageData
import module.logger as logger
import logging
import time
   
################################################
# initialization code
################################################

# 1: Initialize logger!
logger.initLogger()
log = logging.getLogger(__name__)
log.info("Server Started")

# 2. Setup bottle app
app = b.Bottle()
app.install(u.log_to_logger)

# 3. Data setup
user_passwds = u.create_passwd_dict(people)
online_users = {} # Dict[email_id, (user_id, active_at)]

u.SESSION_TIMEOUT_SEC = 600 # seconds
u.AVATAR_IMG_SIZE = (256, 256)  # pixels
u.DEFAULT_IMAGE_QUALITY  = 30  # 30% percent of original

u.ROOT_DIR = osp.dirname(osp.realpath(__file__))
log.info("u.ROOT_DIR = %s", u.ROOT_DIR)
u.TMP_DIR = os.path.join(u.ROOT_DIR, "tmp")
u.STATIC_DIR = os.path.join(u.ROOT_DIR, "static")
u.PEOPLE_DIR = os.path.join(u.ROOT_DIR, "static/image/people")
u.IMAGE_DIR = os.path.join(u.ROOT_DIR, "static/image") 
u.WEB_PEOPLE_DIR = "db"
u.WEB_FREEPIC_DIR = "free"
u.WEB_FILLINS_DIR = "free/fillins"

# assign default pics
u.assign_pics_to_people(people)

################################################


@app.route("/")
@b.view("index") # ./views/index.tpl
def index():
  userid = u.is_logged_in(b.request, online_users)
  if userid is not None:
    b.redirect("/person/{}".format(userid))
  return dict()
                              
@app.post("/login")
@b.view("login") # ./views/login.tpl
def login():
  userid = u.is_logged_in(b.request, online_users)
  if userid is not None:
    b.redirect("/")

  email_id = b.request.forms.get('email_id')
  passwd = b.request.forms.get('passwd')
  person_id = u.is_correct_login(email_id, passwd, user_passwds)
  if person_id:
    online_users[email_id] = (person_id, time.time())
    b.response.set_cookie('email_id', str(email_id))
    b.redirect(f"/person/{person_id}")
  else:
    b.abort(404, "Wrong Email/Passwd!")

@app.get("/login")
def do_login():
  userid = u.is_logged_in(b.request, online_users)
  if userid is None:
    b.redirect("/person/{userid}")
  else:
    b.redirect("/")
                              
@app.get("/logout")
def logout():
  userid = u.is_logged_in(b.request, online_users)
  if userid is not None:
    b.response.set_cookie('email_id', "None")
  b.redirect("/")
                              
@app.route("/signup")
@b.view("signup") # ./views/signup.tpl
def signup():
  b.abort(404, "TODO: Yet to be added!")
                              
@app.route("/person/<person_id:int>")
@b.view("person") # ./views/person.tpl
def person(person_id):
  log.info("try page person_id : %s", person_id)
  userid = u.is_logged_in(b.request, online_users)
  if userid is None:
    b.abort(404, "Not logged in / idle for too long!")

  if person_id in people:
    pd = PageData()
    pd.person = people[person_id]
    pd.relatives = u.fetch_family(pd.person, people)

    pic_dir  = u.person_dir(person_id)
    pd.pics = u.pics_path(pic_dir, pd.person)
    pd.family = u.fetch_family(pd.person, people)
    # pd.profile_pic = u.profile_pic_path(pic_dir, person_id)
    # pd.avatar_pic = u.avatar_pic_path(pic_dir, person_id)
    # (pd.avatar_pic, pd.profile_pic) = u.random_avatar_profile_pic()

    log.info("avatar = %s, profile = %s, pics = %s", pd.avatar_pic, pd.profile_pic, pd.pics)

    return pd.__dict__

  b.abort(404, "Person not found!")
           
@app.route('/query/name/<string>/<randint:int>')
def query_name(string, randint):
    log.info("given string: %s", string)

    result = u.search_name(string, people)

    log.info("matched results: %s", result)

    returnDict = {"rand_int": randint, "data": result}

    b.response.content_type = "application/json"
    return json.dumps(returnDict)

@app.route('/static/<filename:path>')
def send_static(filename):
    return b.static_file(filename, root='./static')                   

@app.route('/backupdb')
def backupdb(filename):
    u.backupdb(people)
    return "TODO"

@app.route("/form")
@b.view("form")  # ./views/form.tpl 
def form():
  return dict()

@app.error(404)
@b.view("error")  # ./views/error.tpl
def error404(error):
  log.error("code = %s, msg = %s", error.status, error.body)
  return dict(msg=error.body + " :(")

#@app.error(500)
#@b.view("error")  # ./views/error.tpl
#def error500(error):
#  return dict(msg="oops !")
#

if __name__ == "__main__":
    app.run(host='localhost', port=8080, debug=True, quiet=True)

