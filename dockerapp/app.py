from flask import Flask, jsonify
from flask_restful import Api, Resource, reqparse, abort, fields, marshal_with


app = Flask(__name__)
api = Api(app)

class index(Resource):
    def get(self):
        return jsonify({"HELLO":"THERE"})

api.add_resource(index,"/")

if __name__ == "__main__":
	app.run(host='0.0.0.0',debug=True)