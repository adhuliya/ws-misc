import flask as f

app = f.Flask(__name__)


@app.route("/")
def index():
    d = {"title": "Hello"}
    return f.render_template("index.html", **d)

@app.route("/favicon.ico")
def favicon():
    return f.send_from_directory("static/images/logo", "favicon.ico")

@app.route("/static/<path:path>")
def staticFiles():
    return f.send_from_directory('static', path)


