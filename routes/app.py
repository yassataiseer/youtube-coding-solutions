
from flask import Flask, render_template,request,session


app = Flask(__name__)


@app.route("/")
def login():
    return render_template("login.html")
@app.route("/signup")
def signup():
    return render_template("signup.html")

app.run( debug=True) 

