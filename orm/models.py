from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import pymysql

host = 'localhost'
user = 'root'
passwd = 'new_password'
database = 'ORMDB'

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'mysql+pymysql://'+user+':'+passwd+'@'+host+"/"+database
app.config['SQLALCHEMY_TRACK_MODIFICATIONS']= False

db = SQLAlchemy(app)

class User(db.Model):
    Username = db.Column(db.String(255), unique=True, nullable=False, primary_key=True)
    Password = db.Column(db.String(255), nullable=False)

class Posts(db.Model):
    Username = db.Column(db.String(255), unique=True, nullable=False, primary_key=True)
    Blog_text = db.Column(db.String(5000),unique=False,nullable=False)

db.create_all()