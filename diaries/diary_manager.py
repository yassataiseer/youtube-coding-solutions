import mysql.connector

class manager:
    def __init__(self,title,content):
        self.title = title
        self.content = content
    def connect():
        db = mysql.connector.connect(
        host = "localhost",
        user = "root",
        passwd = "new_password",
        database = "diary")
        return db
    def write_db(self):
        db = manager.connect()
        mycursor = db.cursor()
        mycursor.execute("INSERT INTO entries (Title,Content) VALUES (%s,%s)",(self.title,self.content))
        db.commit()
        mycursor.close()
        db.close()
        return {"Status":True}
    def delete_db(self):
        db = manager.connect()
        mycursor = db.cursor()
        mycursor.execute("DELETE FROM entries WHERE Title=%s AND Content=%s",(self.title,self.content))
        db.commit()
        mycursor.close()
        db.close()
        return {"Status":True}     


'''
new_entry = manager("Day 1","I made a sandwich today")
print(new_entry.delete_db())'''

