#!/usr/bin/env python3
"""
A bottle app to manage short urls.
module.settings module has most global (if not all) setting parameters.
"""
import module.util as u

# A simple bottle based web app for family tree
import sys
import os
import os.path as osp
import json
import time
import logging

# 1: importing module.settings brings,
#    settings in module/settings.py into effect.
import module.settings

from bottle import Bottle,\
    redirect, abort, request, view,\
    response, static_file

from module.util import is_logged_in, is_correct_login,\
    log_to_logger, delete_short_url, return_json

from module.data import user_map, url_map,\
    PageData, online_users, load_data

   
################################################
# initialization code
################################################

# 2: get logger
log = logging.getLogger(__name__)
log.info("Server Started")

# 3: setup data
load_data()

# 3. Setup bottle app
app = Bottle()
app.install(log_to_logger)

################################################

@app.get("/<string>")
def forward(string):
    url = None
    if string in url_map:
      url = url_map[string]
    if url:
      redirect(url)
    else:
      abort(404, "URL Not Found!")

@app.get("/testit")
def testing():
  return "Test Correct!"

# json request and response only
@app.post("/u/shorten-url")
def shorten_url():
  if request.headers.get("Content-Type") != "application/json":
    return dict(error="Request must be a json object.")

  email_id = is_logged_in(request, online_users)
  if email_id is None:
    redirect("/u/form/login")

  
@app.get("/")
@view("index") # ./views/index.tpl
def index():
  email_id = is_logged_in(request, online_users)
  log.info("good user: %s", email_id)
  if email_id is None:
    redirect("/u/form/login")

  return PageData(user=user_map[email_id], url_map=url_map).__dict__

@app.get("/u/del/<short_url>")
def del_url(short_url):
  email_id = is_logged_in(request, online_users)
  if email_id is None:
    redirect("/u/form/login")

  success = delete_short_url(user_map[email_id], url_map, short_url)

  if success:
    return return_json({'success':success})
  else:
    abort(405, "{success:false}")

@app.get("/u/form/login")
@view("form_login") # ./views/form_login.tpl
def form_login():
  email_id = is_logged_in(request, online_users)
  log.info("login good user: %s", email_id)
  if email_id is not None:
    redirect("/")

  return dict()
                              
@app.post("/u/form/login")
def login():
  email_id = is_logged_in(request, online_users)
  if email_id is not None:
    redirect("/")

  email_id = request.forms.get('email_id')
  passwd = request.forms.get('passwd')
  email_id = is_correct_login(email_id, passwd, user_map)
  if email_id:
    online_users[email_id] = time.time()
    response.set_cookie('email_id', str(email_id), path="/")
    redirect("/")
  else:
    abort(404, "Wrong Email/Passwd!")

@app.get("/u/form/register")
@view("form_register") # ./views/form_register.tpl
def form_register():
  email_id = is_logged_in(request, online_users)
  if email_id is not None:
    redirect("/")

  return dict()
                              
@app.post("/u/form/register")
def register():
  redirect("/")

@app.get("/u/logout")
def logout():
  email_id = is_logged_in(request, online_users)
  if email_id is not None:
    response.set_cookie('email_id', "None")
  redirect("/u/login-form")
                              
@app.error(404)
@view("error") # ./views/error.tpl
def error404(error):
  return dict(msg=error.body)

@app.route('/static/<filename:path>')
def server_static_files(filename):
    return static_file(filename, root='./static')                   


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True, quiet=True)


