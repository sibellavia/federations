from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_talisman import Talisman
import os

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///node.db'
app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY', 'your_secret_key')
db = SQLAlchemy(app)
talisman = Talisman(app)

print("Config and DB setup complete")

# Import routes
from routes import *

if __name__ == "__main__":
    print("Entering main block")
    # Create database and tables
    with app.app_context():
        print("Creating database and tables...")
        db.create_all()
        print("Database and tables created!")
    app.run(ssl_context='adhoc', port=5001)
