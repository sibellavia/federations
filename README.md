## Federations PoC

### Directory Structure

```markdown
federations/
├── api-gateway/
│   └── main.go
├── federation-manager/
│   └── main.go
├── user-management/
│   └── main.go
└── messaging/
    └── main.go
```

### API Documentation

#### Base URL
For the API Gateway:
```
http://localhost:8082
```

### Federation Manager Service

#### Create Federation
**Endpoint:** `/createFederation`  
**Method:** `POST`  
**Description:** Creates a new federation.

**Request Body:**
```json
{
  "name": "Federation Name"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Federation Name",
  "service_catalogue": []
}
```

**Example cURL:**
```bash
curl -X POST http://localhost:8082/createFederation \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Federation1"
      }'
```

#### List Federations
**Endpoint:** `/federations`  
**Method:** `GET`  
**Description:** Lists all federations along with their service catalogues.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Federation Name",
    "service_catalogue": [
      {
        "id": 1,
        "federation_id": 1,
        "name": "Service Name",
        "description": "Service Description"
      }
    ]
  }
]
```

**Example cURL:**
```bash
curl -X GET http://localhost:8082/federations
```

#### Add Service to Federation
**Endpoint:** `/federations/addService`  
**Method:** `POST`  
**Description:** Adds a new service to a federation's service catalogue.

**Request Body:**
```json
{
  "federation_id": 1,
  "name": "Service Name",
  "description": "Service Description"
}
```

**Response:**
```json
{
  "id": 1,
  "federation_id": 1,
  "name": "Service Name",
  "description": "Service Description"
}
```

**Example cURL:**
```bash
curl -X POST http://localhost:8082/federations/addService \
  -H "Content-Type: application/json" \
  -d '{
        "federation_id": 1,
        "name": "Service1",
        "description": "Description of Service1"
      }'
```

#### Get Services of Federation
**Endpoint:** `/federations/getServices`  
**Method:** `GET`  
**Description:** Retrieves the services for a specific federation.

**Query Parameters:**
- `federation_id` (required): The ID of the federation.

**Response:**
```json
[
  {
    "id": 1,
    "federation_id": 1,
    "name": "Service Name",
    "description": "Service Description"
  }
]
```

**Example cURL:**
```bash
curl -X GET "http://localhost:8082/federations/getServices?federation_id=1"
```

### User Management Service

#### Register User
**Endpoint:** `/register`  
**Method:** `POST`  
**Description:** Registers a new user.

**Request Body:**
```json
{
  "username": "user1",
  "password": "password123",
  "role": "user"
}
```

**Response:**
```json
{
  "status": "registered"
}
```

**Example cURL:**
```bash
curl -X POST http://localhost:8082/register \
  -H "Content-Type: application/json" \
  -d '{
        "username": "user1",
        "password": "password123",
        "role": "user"
      }'
```

#### Login User
**Endpoint:** `/login`  
**Method:** `POST`  
**Description:** Logs in a user.

**Request Body:**
```json
{
  "username": "user1",
  "password": "password123"
}
```

**Response:**
```json
{
  "status": "login successful",
  "token": "oauth2_access_token"
}
```

**Example cURL:**
```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{
        "username": "user1",
        "password": "password123"
      }'
```

### Messaging Service

#### Send Message
**Endpoint:** `/sendMessage`  
**Method:** `POST`  
**Description:** Sends a message between federations.

**Request Body:**
```json
{
  "from": "user1@fed-1",
  "to": "user2@fed-2",
  "content": "Hello from Federation 1!"
}
```

**Response:**
```json
{
  "status": "message sent"
}
```

**Example cURL:**
```bash
curl -X POST http://localhost:8082/sendMessage \
  -H "Content-Type: application/json" \
  -d '{
        "from": "user1@fed-1",
        "to": "user2@fed-2",
        "content": "Hello from Federation 1!"
      }'
```

#### Get Messages
**Endpoint:** `/getMessages`  
**Method:** `GET`  
**Description:** Retrieves all messages.

**Response:**
```json
[
  {
    "from": "user1@fed-1",
    "to": "user2@fed-2",
    "content": "Hello from Federation 1!"
  }
]
```

**Example cURL:**
```bash
curl -X GET http://localhost:8082/getMessages
```
