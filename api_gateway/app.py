from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_talisman import Talisman
import os

app = Flask(__name__)
app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY', 'your_secret_key')
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///gateway.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)
talisman = Talisman(app)

class Node(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    node_name = db.Column(db.String(256), unique=True)
    node_url = db.Column(db.String(256))
    auth_token = db.Column(db.String(256))

print("Creating the database and tables if not exists...")
with app.app_context():
    db.create_all()
    print("Database and tables created!")

@app.route('/api/nodes/register', methods=['POST'])
def register_node():
    data = request.get_json()
    new_node = Node(node_name=data['node_name'], node_url=data['node_url'], auth_token=data['auth_token'])
    db.session.add(new_node)
    db.session.commit()
    return jsonify({"status": "Node Registered"})

@app.route('/api/gateway/forward', methods=['POST'])
def forward_request():
    data = request.get_json()
    target_node_name = data['target_node']
    message = data['message']
    
    node = Node.query.filter_by(node_name=target_node_name).first()
    if not node:
        return jsonify({"status": "Node not found"}), 404
    
    response = request.post(node.node_url, json={"message": message})
    return jsonify({"status": "Forwarded", "response": response.json()})

if __name__ == "__main__":
    app.run(ssl_context='adhoc', port=5000)