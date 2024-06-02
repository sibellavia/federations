from app import db

class RequestLog(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    message = db.Column(db.String(256))
    timestamp = db.Column(db.DateTime, default=db.func.current_timestamp())
