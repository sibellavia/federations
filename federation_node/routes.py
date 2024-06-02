from app import app, db
from flask import request, jsonify
from models import RequestLog
import requests

@app.route('/api/receive', methods=['POST'])
def receive_request():
    data = request.get_json()
    # Log the request
    new_log = RequestLog(message=data.get('message'))
    db.session.add(new_log)
    db.session.commit()
    return jsonify({"status": "Received"})

@app.route('/api/send', methods=['POST'])
def send_request():
    target_node = request.json.get('target_node')
    message = request.json.get('message')
    # Send request to the target node
    response = requests.post(target_node, json={"message": message})
    return jsonify({"status": "Sent", "response": response.json()})
