from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_talisman import Talisman
import os
import requests

app = Flask(__name__)
app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY', 'your_secret_key')
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///node.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)
talisman = Talisman(app)

class RequestLog(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    message = db.Column(db.String(256))
    timestamp = db.Column(db.DateTime, default=db.func.current_timestamp())

print("Creating the database and tables if not exists...")
with app.app_context():
    db.create_all()
    print("Database and tables created!")

API_GATEWAY_URL = 'https://127.0.0.1:5000/api/nodes/register'

def register_with_gateway(node_name, node_url, auth_token):
    registration_data = {
        "node_name": node_name,
        "node_url": node_url,
        "auth_token": auth_token
    }
    try:
        response = requests.post(API_GATEWAY_URL, json=registration_data, verify=False)
        if response.status_code == 200:
            print(f"Successfully registered {node_name} with API Gateway")
            return True
        else:
            print(f"Failed to register with API Gateway: {response.status_code} {response.text}")
            return False
    except Exception as e:
        print(f"Error registering with API Gateway: {e}")
        return False

@app.route('/api/receive', methods=['POST'])
def receive_request():
    data = request.get_json()
    new_log = RequestLog(message=data.get('message'))
    db.session.add(new_log)
    db.session.commit()
    return jsonify({"status": "Received"})

@app.route('/api/send', methods=['POST'])
def send_request():
    target_node = request.json.get('target_node')
    message = request.json.get('message')
    response = requests.post(target_node, json={"message": message})
    return jsonify({"status": "Sent", "response": response.json()})

@app.route('/api/register_node', methods=['POST'])
def register_node_endpoint():
    data = request.get_json()
    node_name = data.get('node_name')
    node_url = data.get('node_url')
    auth_token = data.get('auth_token')
    if register_with_gateway(node_name, node_url, auth_token):
        return jsonify({"status": "Registered"})
    else:
        return jsonify({"status": "Failed"}), 500

if __name__ == "__main__":
    app.run(ssl_context='adhoc', port=5001)
