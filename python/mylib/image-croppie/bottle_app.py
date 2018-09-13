#!/usr/bin/env python3

import bottle as b

app = b.Bottle()

@app.get("/")
@b.view("index")
def index():
    return dict()

@app.route('/static/<filename:path>')
def send_static(filename):
    return b.static_file(filename, root='./static')                   

@app.post('/path/to/upload')
def upload_image():
  img = b.request.forms.get('img_blob', "")
  img_type = b.request.forms.get('img_type', "")
  print(len(img), img_type)
  return ""

if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5001", debug=True)
