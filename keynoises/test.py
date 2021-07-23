

import YassaAlchemy 

db = YassaAlchemy.table(host="localhost",user="root",passwd="new_password",database="yassadb")

class Users:

    def create(db):
        db.connect()
        db.mk_table("Users")
        db.mk_column("Username","VARCHAR(5000)")
        db.mk_column("Password","VARCHAR(5000)")
        db.publish_db()

Users.create(db)



