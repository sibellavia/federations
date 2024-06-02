# Federations PoC (Proof of Concept)

## Project Structure

```markdown
federations/
│
├── api_gateway/
│   ├── app.py
│   ├── __init__.py
│   └── requirements.txt
│
├── federation_node/
│   ├── app.py
│   ├── __init__.py
│   └── requirements.txt
│
├── integration_layer/
│   ├── app.py
│   ├── __init__.py
│   └── requirements.txt
│
├── docker-compose.yml
└── README.md
```

## Architecture.

A PoC that allows different federated instances to communicate with each other via API. Special focus on the integration and API part without worrying too much about a UI. 

Main components:
- Federation Nodes
- API Gateway
- Integration Layer
- Authentication and Security

### Federation Nodes.

Each node represents a federated instance. Nodes must be able to receive, process, and respond to requests.

### API Gateway.

The API Gateway serves as the central point for routing requests between federated nodes.

### Integration Layer.

The Integration Layer manages the integrations between different cloud infrastructures.

### Authentication and Security.

Authentication and Security ensures that only authorized nodes can communicate with each other.

### Diagram

```lua
+----------------------+          +----------------------+          +----------------------+
|  Federation Node A   | <------> |      API Gateway     | <------> |  Federation Node B   |
+----------------------+          +----------------------+          +----------------------+
        ^  ^                                                        ^  ^
        |  |                                                        |  |
        |  +-------------------+                +-------------------+  |
        |                      |                |                      |
        |   +--------------------+            +--------------------+   |
        |   | Integration Layer  |            | Integration Layer  |   |
        |   +--------------------+            +--------------------+   |
        |                                                             |
        +-------------------------------------------------------------+
```

## Tech Stack

- Framework: Flask
- Database: SQLite
- Autenticazione: pyjwt
- HTTPS: Flask-Talisman
and Docker.

## API Documentation 

### Overview

The PoC consists of three main components:

1. **Federation Node**: Each node represents a federated instance.
2. **API Gateway**: Central gateway for routing requests between federation nodes.
3. **Integration Layer**: Handles integration with different cloud providers.

### Base URL

Each component runs on a different port. The base URLs for each component are as follows:

- **Federation Node**: `http://localhost:5001`
- **API Gateway**: `http://localhost:5000`
- **Integration Layer**: `http://localhost:5002`

### Federation Node API

#### 1. Receive Request

**Endpoint**: `/api/receive`

**Method**: `POST`

**Description**: Receives a request from another federation node.

**Request Body**:
```json
{
  "message": "Hello from Node <node_name>"
}
```

**Response**:
```json
{
  "status": "Received"
}
```

#### 2. Send Request

**Endpoint**: `/api/send`

**Method**: `POST`

**Description**: Sends a request to another federation node.

**Request Body**:
```json
{
  "target_node": "http://node_b/api/receive",
  "message": "Hello from Node <node_name>"
}
```

**Response**:
```json
{
  "status": "Sent",
  "response": "Received"
}
```

### API Gateway

#### 1. Register Node

**Endpoint**: `/api/nodes/register`

**Method**: `POST`

**Description**: Registers a new federation node.

**Request Body**:
```json
{
  "node_name": "Node A",
  "node_url": "http://node_a/api/receive",
  "auth_token": "node_a_token"
}
```

**Response**:
```json
{
  "status": "Node Registered"
}
```

#### 2. Forward Request

**Endpoint**: `/api/gateway/forward`

**Method**: `POST`

**Description**: Forwards a request to a specified federation node.

**Request Body**:
```json
{
  "target_node": "Node B",
  "message": "Hello from Gateway"
}
```

**Response**:
```json
{
  "status": "Forwarded",
  "response": "Received"
}
```

### Integration Layer

#### 1. Get Cloud Resources

**Endpoint**: `/api/cloud/resources`

**Method**: `GET`

**Description**: Returns a list of available resources from a specified cloud provider.

**Response**:
```json
{
  "resources": [
    "VM1",
    "VM2",
    "Storage1"
  ]
}
```

#### 2. Execute Cloud Action

**Endpoint**: `/api/cloud/action`

**Method**: `POST`

**Description**: Executes a specified action on a cloud resource.

**Request Body**:
```json
{
  "action": "start_vm",
  "resource_id": "VM1"
}
```

**Response**:
```json
{
  "status": "Action Executed"
}
```

### Authentication and Security

All API endpoints require authentication using JSON Web Tokens (JWT). Include the JWT in the `Authorization` header of each request:

**Authorization Header**:
```
Authorization: Bearer <your_jwt_token>
```

### Example Workflow

1. **Register a Node**: A federation node registers itself with the API Gateway.
    - **Request**:
        ```bash
        curl -X POST http://localhost:5000/api/nodes/register -H "Content-Type: application/json" -d '{
          "node_name": "Node A",
          "node_url": "http://node_a/api/receive",
          "auth_token": "node_a_token"
        }'
        ```

2. **Send a Request**: A federation node sends a request to another node.
    - **Request**:
        ```bash
        curl -X POST http://localhost:5001/api/send -H "Content-Type: application/json" -d '{
          "target_node": "http://node_b/api/receive",
          "message": "Hello from Node"
        }'
        ```

3. **Forward a Request**: The API Gateway forwards a request to a specified node.
    - **Request**:
        ```bash
        curl -X POST http://localhost:5000/api/gateway/forward -H "Content-Type: application/json" -d '{
          "target_node": "Node B",
          "message": "Hello from Gateway"
        }'
        ```

4. **Get Cloud Resources**: Retrieve the list of resources from a cloud provider.
    - **Request**:
        ```bash
        curl -X GET http://localhost:5002/api/cloud/resources
        ```

5. **Execute Cloud Action**: Perform an action on a cloud resource.
    - **Request**:
        ```bash
        curl -X POST http://localhost:5002/api/cloud/action -H "Content-Type: application/json" -d '{
          "action": "start_vm",
          "resource_id": "VM1"
        }'
        ```