#!/usr/bin/env python3

# uses jquery library for ajax calls

import bottle as b
from datetime import datetime as dt
import json
from json import dumps

globalvar = 0

app = b.Bottle()

page = """
<!DOCTYPE html><html lang="en"><head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

</head><body>

<button id="btn1" type="button">Send Json Method 1</button>
<p id="display1"></p>

<button id="btn2" type="button">Send Json Method 2</button>
<p id="display2"></p>

<script src="static/js/jquery.min.js"></script>

<script>
var age1 = 10;
var age2 = 10;

$("#btn1").on('click', function () {
  $.ajax({
    url: '/receive_json/method1',
    type: 'post',
    dataType: 'json',
    contentType: 'application/json',
    success: function (data) {
      $('#display1').html(data.msg);
    },
    data: JSON.stringify({age1: age1})
  });
  age1 += 1;
});

$("#btn2").on('click', function () {
  $.ajax({
    url: '/receive_json/method2',
    type: 'post',
    dataType: 'json',
    contentType: 'application/json',
    success: function (data) {
      $('#display2').html(data.msg);
    },
    data: JSON.stringify({age2: age2})
  });
  age2 += 1;
});

</script>

</body></html>
"""

global_var1 = 0
global_var2 = 0

@app.route("/")
def index():
    return page

@app.route('/static/<filename:path>')
def send_static(filename):
    return b.static_file(filename, root='./static')                   

@app.post("/receive_json/method1")
def receive_json1():
    # when Content-Type header is `application/json`
    # only requests smaller than MEMFILE_MAX are processed
    # to void memory exhaustion
    global global_var1
    global_var1 += 1
    data = b.request.json
    print("{}: {}".format(type(data), data))
    return return_json(dict(msg="Method1 data: {}".format(global_var1)))

@app.post("/receive_json/method2")
def receive_json2():
    # when Content-Type header is `application/json`
    global global_var2
    global_var2 += 1
    data = json.load(b.request.body)
    print("{}: {}".format(type(data), data))
    return return_json(dict(msg="Method2 data: {}".format(global_var2)))

def return_json(data):
    b.response.content_type = "application/json"
    return dumps(data)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5000", debug=True)
    print("Started as stand alone program")


