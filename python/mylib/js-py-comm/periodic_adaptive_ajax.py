#!/usr/bin/env python3
# A single file application.

import bottle as b
import json
import time

app = b.Bottle()

# when using it, store this javascript in a separate file,
# may be call it async.js
javascript = """
// Functions to facilitate requesting data from server,
// and coping with the delays in communication by waiting.

// Request server for data. (Json mostly)
function httpGetAsync(theUrl, callback) {
  var xmlHttp = new XMLHttpRequest();

  xmlHttp.onreadystatechange = function() {
    if (xmlHttp.readyState === 4 && xmlHttp.status == 200) {
      callback(xmlHttp.responseText, true);
    } else if (xmlHttp.readyState === 4) {
      // when server doesn't respond in time or there is an error
      callback(undefined, false);
    }
  }

  xmlHttp.open("GET", theUrl, true); // true for asyncronous
  xmlHttp.send(null);
}

// Request server for data. (Json mostly)
function httpGetAsyncObj(theUrl, callback, callbackObj) {
  var xmlHttp = new XMLHttpRequest();

  xmlHttp.onreadystatechange = function() {
    if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
      callback.apply(callbackObj, [xmlHttp.responseText, true]);
    } else if (xmlHttp.readyState === 4) {
      // when server doesn't respond in time or there is an error
      callback.apply(callbackObj, [undefined, false]);
    }
  }

  xmlHttp.open("GET", theUrl, true); // true for asyncronous
  xmlHttp.send(null);
}

function modifyInterval(obj, success) {
  // exponentially shrinks/grows waitTime
  var curWaitCopy = obj.curWait;
  if (success) {
    if (obj.periodic === true) {
      obj.upCounter += 1;
      // recover slowly, fail quickly.
      if (obj.upCounter >= obj.maxCount) {
        if (obj.curWait !== obj.minWait) {
          var waitTime = obj.curWait >> 1;
          if (waitTime < obj.minWait) {
            obj.curWait = minWait;
          } else {
            obj.curWait = waitTime;
          }
        }
        obj.upCounter = 0;
      }
    }
  } else {
    // recover slowly, fail quickly.
    if (obj.curWait !== obj.maxWait) {
      var waitTime = obj.curWait << 1;
      if (waitTime > obj.maxWait) {
        obj.curWait = maxWait;
      } else {
        obj.curWait = waitTime;
      }
    }
  }

  if (obj.periodic === true) {
    if (curWaitCopy !== obj.curWait) {
      window.clearInterval(obj.invokeId);
      obj.invokeId = window.setInterval(obj.sendRequest(obj), obj.curWait);
    }
  } else {
    if (!success) {
      // window.clearTimeout(obj.invokeId);
      obj.invokeId = window.setTimeout(obj.sendRequest(obj), obj.curWait);
    }
  }
}

// Example use of above two functions.
// START: Object that keeps track of repeated server communication
var onetimeRequestObj1 = {
  url: "http://localhost:5000/name",
  periodic: false,
  curWait: 1 << 11, // 1 << 11 = 2048 ~ 2 second
  minWait: 1 << 11, // 1 << 11 = 2048 ~ 2 second
  maxWait: 1 << 15, // 1 << 15 = 32768 ~ 33 second

  // the var to hold the window.setInterval id/window.setTimeout id.
  invokeId: undefined,

  upCounter: 0,
  maxCount: 2,

  sendRequest: function (obj) {
    return function() {
      //console.log("SENDING REQUEST");
      httpGetAsyncObj(obj.url, obj.callBack, obj);
      // for demonstration and testing purpose
      document.getElementById("timeout").innerHTML = obj.curWait;
    }
  },

  callBack: function (data, success) {
    // This method is executed whenever server sends data or
    // the request fails due to some reason (success === false).
    modifyInterval(this, success);
    if (success) {
      var elementId = "thename";
      var obj = JSON.parse(data);
      document.getElementById(elementId).innerHTML = obj.name;
    }
  },

  start: function () {
    this.invokeId = window.setTimeout(this.sendRequest(this), this.curWait);
  }
};

// Example use of above functions, for ONE TIME communication
// keep trying until success
// START: Object that keeps track of repeated server communication
var periodicRequestObj1 = {
  url: "http://localhost:5000/time",
  periodic: true,
  curWait: 1 << 11, // 1 << 11 = 2048 ~ 2 second
  minWait: 1 << 11, // 1 << 11 = 2048 ~ 2 second
  maxWait: 1 << 15, // 1 << 15 = 32768 ~ 33 second

  // the var to hold the window.setInterval id.
  invokeId: undefined,

  upCounter: 0,
  maxCount: 2,

  sendRequest: function (obj) {
    return function() {
      //console.log("SENDING REQUEST");
      httpGetAsyncObj(obj.url, obj.callBack, obj);
      // for demonstration and testing purpose
      document.getElementById("interval").innerHTML = obj.curWait;
    }
  },

  callBack: function (data, success) {
    // This method is executed whenever server sends data or
    // the request fails due to some reason (success === false).
    modifyInterval(this, success);
    if (success) {
      var elementId = "timer";
      var obj = JSON.parse(data);
      document.getElementById(elementId).innerHTML = obj.time;
    }
  },

  start: function () {
    this.invokeId = window.setInterval(this.sendRequest(this), this.curWait);
  }
};

// for one time non-repeating requests
// irrespective of if it fails or succeeds
// rand_int is used to map requests to response correctly
// every new request resets rand_int,
// effectively making any previous request garbage.
var throwAwayRequestObj = {
  rand_int : 0,
  string: "",
  url: "/query/name/",

  callBack: function (data, success) {
    if (success) {
      var elementId = "data-test";
      var obj = JSON.parse(data);
      console.log("TRYING :", this.rand_int, obj.rand_int, obj.data);
      if (this.rand_int == obj.rand_int) {
        console.log("SUCCESS: ", this.rand_int, obj.rand_int, obj.data);
        // race condition possible here.
        this.rand_int = 0; // reset
        // TODO : the action to be taken
        document.getElementById(elementId).innerHTML = obj.data;
      }
    } else {
      console.log("not same rand_int");
      this.rand_int = 0;
    }
  },

  sendRequest: function (string) {
    this.string = string;
    // new random variable for new request
    this.rand_int = Math.floor(Math.random() * 1000000000) + 1;
    var query_url = this.url + this.string + "/" + this.rand_int;
    httpGetAsyncObj(query_url, this.callBack, this);
    console.log("SENT: ", this.rand_int, this.string);
  }
};

//Call the start() method from the page using this file.
//called from `/template/timer.tmpl`
//periodicRequestObj1.start();

//For one time invocation:
//id1 = window.setTimeout(callHttpGetAsync, 1000);
//To Cancel it:
//window.clearTimeout(id1);
// END  : Object that keeps track of repeated server communication
"""

page = """
<!DOCTYPE html>
<html>
<body>
A page demonstrating adaptive periodic communication in presence of
communication failure.
<br/> <br/>
Value from server: <span id="timer"></span>
<br/>
Fetch Interval: <span id="interval"></span>
<br/> <hr/>
Request Timeout: <span id="timeout"></span>
<br/>
Name: <span id="thename">This should change.</span>

<br/>
<br/>
Data Test: <span id="data-test"></span>
<script>
{javascript}
</script>
<script>
periodicRequestObj1.start();
onetimeRequestObj1.start();
throwAwayRequestObj.sendRequest("anu"); 
throwAwayRequestObj.sendRequest("au"); 
throwAwayRequestObj.sendRequest("a"); 
throwAwayRequestObj.sendRequest("sdf"); 
throwAwayRequestObj.sendRequest("sf"); 
throwAwayRequestObj.sendRequest("s"); 
throwAwayRequestObj.sendRequest("w"); 
</script>
</body>
</html>
""".format(javascript=javascript)

errTimes = 0

@app.get("/")
def index():
  return page

@app.get("/name")
def name():
  """deliberately fail the request for first 5 times"""
  global errTimes
  if errTimes < 5:
    errTimes += 1
    b.abort(404, "smthing")
  else:
    return return_json({'name': "Anshuman Dhuliya"})

@app.get("/time")
def sendtime():
    return return_json({'time': time.time()})

@app.get("/query/name/<name>/<randint:int>")
def send_data(name, randint):
    return return_json({'rand_int': randint, 'data': 'this is data for ' + name})

def return_json(d):
    b.response.content_type = "application/json"
    return json.dumps(d)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)


