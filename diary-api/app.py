from entry_manager import manager
from flask import Flask, jsonify
from flask_restful import Api, Resource, reqparse, abort, fields, marshal_with

app = Flask(__name__)
api = Api(app)

class mk_diaries(Resource):
    def get(self,Title,Content):
        new_entry = manager(Title,Content)
        status = new_entry.write_db()
        return jsonify(status)
class delete_diary(Resource):
    def get(self,Title,Content):
        del_entry = manager(Title,Content)
        status = del_entry.delete_order()
        return jsonify(status)



api.add_resource(mk_diaries,"/mk_diaries/<string:Title>/<string:Content>")
api.add_resource(delete_diary,"/delete_diary/<string:Title>/<string:Content>")
if __name__ == "__main__":
	app.run(port=5000 ,debug=True)


