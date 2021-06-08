#main.py
from models import User,db

exists = bool(db.session.query(User).filter_by(Username="Mark Taiseer").first())
print(exists)