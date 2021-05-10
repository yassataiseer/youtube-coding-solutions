from flask import Flask, jsonify
from flask_restful import Api, Resource, reqparse, abort, fields, marshal_with
from Diary.entry_template import entry_template
app = Flask(__name__)
api = Api(app)

app.register_blueprint(entry_template,url_prefix="/Diary")




if __name__ == "__main__":
	app.run(port=5000 ,debug=True)


