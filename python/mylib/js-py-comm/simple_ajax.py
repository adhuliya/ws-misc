#!/usr/bin/env python3
import bottle as b
from datetime import datetime as dt
from json import dumps

globalvar = 0

app = b.Bottle()

page = """
<!DOCTYPE html>
<html>

<head>
</head>

<body>
<h2>Json and Bottlepy communication</h2>

<p id="date"></p>
<p id="counter"></p>
<p id="square"></p>

<script>
function httpGetAsync(theUrl, callback) {
  var xmlHttp = new XMLHttpRequest();

  xmlHttp.onreadystatechange = function() {
    if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
      callback(xmlHttp.responseText);
    }
  }

  xmlHttp.open("GET", theUrl, true); // true for asyncronous
  xmlHttp.send(null);
}

function setDate(dateStr) {
  document.getElementById("date").innerHTML = dateStr;
}

function setCounter(counterJson) {
  var obj = JSON.parse(counterJson);
  document.getElementById("counter").innerHTML = obj.counter;
  document.getElementById("square").innerHTML = obj.square;
}

function callHttpGetAsync() {
  httpGetAsync("http://localhost:5000/date", setDate);
}

function callHttpGetAsyncWithCounter() {
  httpGetAsync("http://localhost:5000/json", setCounter);
}

var intId1 = window.setInterval(callHttpGetAsync, 1000);
var intId2 = window.setInterval(callHttpGetAsyncWithCounter, 100);

//to stop the interval
//window.clearInterval(intId1);
//window.clearInterval(intId1);
/*
//For one time invocation:
id1 = window.setTimeout(callHttpGetAsync, 1000);
//To Cancel it:
window.clearTimeout(id1);
*/

</script>
</body>
</html>
"""

@app.route("/")
def index():
    return page

@app.route("/date")
def dateString():
    return str(dt.now())

@app.route("/json")
def counterJson():
    print("someone called me")
    global globalvar
    d = {"counter": globalvar, "square": globalvar * globalvar}
    globalvar += 1
    return return_json(d)

def return_json(d):
    b.response.content_type = "application/json"
    return dumps(d)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5000", debug=True)
    print("Started as stand alone program")


