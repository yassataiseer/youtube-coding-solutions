from flask import Flask,jsonify
from flask_restful import Api, Resource
from flask_httpauth import HTTPBasicAuth

app = Flask(__name__)
api = Api(app)
auth = HTTPBasicAuth()
USER_DATA = {
   "Username": "password"
}

@auth.verify_password
def verify(username, password):
   if not (username and password):
       return False
   return USER_DATA.get(username) == password

class endpoint(Resource):
   @auth.login_required
   def get(self):
       return jsonify({"status":True})

api.add_resource(endpoint,"/endpoint")


if __name__=="__main__":
   app.run(port=5000,debug=True)