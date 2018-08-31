
# A very simple Bottle Hello World app for you to get started with...
import sys
# sys.path.append(".")

from bottle import default_app, route, run, static_file, error, view, abort, redirect, request
from module.person import Person
from module.people import people
import module.util as util

passwds = util.create_passwd_dict(people)


@route("/")
@view("index") # ./views/index.tpl
def index():
    return dict()
                              
@route("/login", method="POST")
@view("login") # ./views/login.tpl
def login():
    email_id = request.forms.get('email_id')
    passwd = request.forms.get('passwd')
    if email_id in passwds:
      val = passwds[email_id] 
      if val[1] == passwd:
        redirect(f"/person/{val[0]}")
    else:
      abort(404, "Wrong Email/Passwd!")
                              
@route("/signup")
@view("signup") # ./views/signup.tpl
def signup():
    return "TODO"
                              
@route("/person/<person_id:int>")
@view("person") # ./views/person.tpl
def person(person_id):
    if person_id in people:
        p = people[person_id]
        return p.__dict__
    abort(404, "Person not found")
           
@route('/static/<filename:path>')
def send_static(filename):
    return static_file(filename, root='./static')                   

@route('/backupdb')
def backupd(filename):
    util.backupdb(people)
    return "TODO"

@error(404)
@view("error")  # ./views/error.tpl
def error404(error):
  return dict(msg="Not available :(")

@error(500)
@view("error")  # ./views/error.tpl
def error500(error):
  return dict(msg="oops !")

application = default_app()  

if __name__ == "__main__":
    run(host='localhost', port=8080)

