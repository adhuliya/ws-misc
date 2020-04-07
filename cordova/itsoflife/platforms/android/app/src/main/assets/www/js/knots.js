// A custom js library for knots

function myopen(elementA) {
  var open = cordova.InAppBrowser.open;
  var ref = open(elementA.href, "_blank", "location=no");
}
