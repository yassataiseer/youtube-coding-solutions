import mysql.connector

db = mysql.connector.connect(
    host = "localhost",
    user = "root",
    passwd = "new_password",
    database = "diary"
)
mycursor = db.cursor()
mycursor.execute("CREATE TABLE entries(Title VARCHAR(50), Content VARCHAR(5000))")

