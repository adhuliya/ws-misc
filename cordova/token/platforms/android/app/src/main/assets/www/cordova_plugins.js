cordova.define('cordova/plugin_list', function(require, exports, module) {
  module.exports = [
    {
      "id": "cordova-plugin-inappbrowser.inappbrowser",
      "file": "plugins/cordova-plugin-inappbrowser/www/inappbrowser.js",
      "pluginId": "cordova-plugin-inappbrowser",
      "clobbers": [
        "cordova.InAppBrowser.open",
        "window.open"
      ]
    },
    {
      "id": "cordova-plugin-sms.SMS",
      "file": "plugins/cordova-plugin-sms/www/SMS.js",
      "pluginId": "cordova-plugin-sms",
      "clobbers": [
        "window.SMS"
      ]
    }
  ];
  module.exports.metadata = {
    "cordova-plugin-whitelist": "1.3.4",
    "cordova-plugin-inappbrowser": "3.2.0",
    "cordova-plugin-sms": "1.0.5"
  };
});