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
  // exponentially shrinks/grows interval
  var curIntervalCopy = obj.curInterval;
  if (success) {
    obj.upCounter += 1;
    // recover slowly, fail quickly.
    if (obj.upCounter >= obj.maxCount) {
      if (obj.curInterval !== obj.minInterval) {
        var interval = obj.curInterval >> 1;
        if (interval < obj.minInterval) {
          obj.curInterval = minInterval;
        } else {
          obj.curInterval = interval;
        }
      }
      obj.upCounter = 0;
    }
  } else {
    // recover slowly, fail quickly.
    if (obj.curInterval !== obj.maxInterval) {
      var interval = obj.curInterval << 1;
      if (interval > obj.maxInterval) {
        obj.curInterval = maxInterval;
      } else {
        obj.curInterval = interval;
      }
    }
  }

  if (curIntervalCopy !== obj.curInterval) {
    window.clearInterval(obj.intervalId);
    obj.intervalId = window.setInterval(obj.sendRequest(obj), obj.curInterval);
  }
}

// Example use of above two functions.
// START: Object that keeps track of repeated server communication
var updateObj01 = {
  url: "http://localhost:9090/time",
  curInterval: 1 << 11, // 1 << 11 = 2048 ~ 2 second
  minInterval: 1 << 11, // 1 << 11 = 2048 ~ 2 second
  maxInterval: 1 << 15, // 1 << 15 = 32768 ~ 33 second

  // the var to hold the window.setInterval id.
  intervalId: undefined,

  upCounter: 0,
  maxCount: 2,

  sendRequest: function (obj) {
    return function() {
      //console.log("SENDING REQUEST");
      httpGetAsyncObj(obj.url, obj.callBack, obj);
      // for demonstration and testing purpose
      document.getElementById("interval").innerHTML = obj.curInterval;
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
    this.intervalId = window.setInterval(this.sendRequest(this), this.curInterval);
  }
};

//Call the start() method from the page using this file.
//called from `/template/timer.tmpl`
//updateObj01.start();

// END  : Object that keeps track of repeated server communication


