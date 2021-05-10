from flask import Flask, jsonify,Blueprint
from flask_restful import Api,Resource
from Diary.entry_manager import manager

entry_template = Blueprint("entry_template",__name__)


api = Api(entry_template)

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


