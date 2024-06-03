## Federations PoC

## Basic Features

1. **Federation Management:**
   - **Create Federation:** Allows the creation of new federation instances.
   - **List Federations:** Retrieves a list of all existing federation instances.

2. **User Management:**
   - **Register User:** Allows users to register within a federation.
   - **Login User:** Authenticates users and provides them with an access token.

3. **Messaging:**
   - **Send Message:** Enables users to send messages between federations.
   - **Get Messages:** Retrieves all messages for the federation.

## Basic Workflows

1. **Creating a New Federation:**
   - **Step 1:** The admin sends a request to the API Gateway to create a new federation using the `/createFederation` endpoint.
   - **Step 2:** The API Gateway forwards the request to the Federation Manager service.
   - **Step 3:** The Federation Manager creates a new federation instance and stores its details.
   - **Step 4:** The Federation Manager returns the federation ID and status to the API Gateway, which then forwards the response to the admin.

2. **Listing Existing Federations:**
   - **Step 1:** A user sends a request to the API Gateway to list all federations using the `/federations` endpoint.
   - **Step 2:** The API Gateway forwards the request to the Federation Manager service.
   - **Step 3:** The Federation Manager retrieves the list of federations and sends it back to the API Gateway, which then forwards the response to the user.

3. **Registering a New User:**
   - **Step 1:** A user sends a request to the API Gateway to register using the `/register` endpoint.
   - **Step 2:** The API Gateway forwards the request to the User Management service.
   - **Step 3:** The User Management service registers the user and stores their details.
   - **Step 4:** The User Management service returns the user ID and status to the API Gateway, which then forwards the response to the user.

4. **Logging in a User:**
   - **Step 1:** A user sends a request to the API Gateway to log in using the `/login` endpoint.
   - **Step 2:** The API Gateway forwards the request to the User Management service.
   - **Step 3:** The User Management service authenticates the user and generates an access token.
   - **Step 4:** The User Management service returns the access token and status to the API Gateway, which then forwards the response to the user.

5. **Sending a Message:**
   - **Step 1:** A user sends a request to the API Gateway to send a message using the `/sendMessage` endpoint.
   - **Step 2:** The API Gateway forwards the request to the Messaging service.
   - **Step 3:** The Messaging service processes the message and stores it.
   - **Step 4:** The Messaging service returns the message ID and status to the API Gateway, which then forwards the response to the user.

6. **Retrieving Messages:**
   - **Step 1:** A user sends a request to the API Gateway to retrieve messages using the `/getMessages` endpoint.
   - **Step 2:** The API Gateway forwards the request to the Messaging service.
   - **Step 3:** The Messaging service retrieves all messages and sends them back to the API Gateway, which then forwards the response to the user.

Sure! Below is the detailed documentation for the APIs included in your PoC. This documentation will cover endpoints for federation management, user management, and messaging services.

## API Documentation

### Overview

The APIs are divided into three main categories:
1. Federation Management
2. User Management
3. Messaging Service

### Base URL
The base URL for the API Gateway is `http://localhost:8080`.

### 1. Federation Management

#### 1.1 Create Federation
**Endpoint:** `/createFederation`  
**Method:** `POST`  
**Description:** Creates a new federation instance.

**Request:**
```json
{
  "name": "string",
  "admin": {
    "username": "string",
    "password": "string"
  }
}
```

**Response:**
```json
{
  "federationId": "string",
  "status": "string"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/createFederation \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Federation1",
        "admin": {
          "username": "admin",
          "password": "admin123"
        }
      }'
```

**Example Response:**
```json
{
  "federationId": "fed-1",
  "status": "created"
}
```

#### 1.2 List Federations
**Endpoint:** `/federations`  
**Method:** `GET`  
**Description:** Lists all existing federation instances.

**Response:**
```json
{
  "federations": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}
```

**Example Request:**
```bash
curl -X GET http://localhost:8080/federations
```

**Example Response:**
```json
{
  "federations": [
    {
      "id": "fed-1",
      "name": "Federation1"
    }
  ]
}
```

### 2. User Management

#### 2.1 Register User
**Endpoint:** `/register`  
**Method:** `POST`  
**Description:** Registers a new user within a federation.

**Request:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "userId": "string",
  "status": "string"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
        "username": "user1",
        "password": "password123"
      }'
```

**Example Response:**
```json
{
  "userId": "user-1",
  "status": "registered"
}
```

#### 2.2 Login User
**Endpoint:** `/login`  
**Method:** `POST`  
**Description:** Authenticates a user and returns an access token.

**Request:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "token": "string",
  "status": "string"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
        "username": "user1",
        "password": "password123"
      }'
```

**Example Response:**
```json
{
  "token": "abcdef123456",
  "status": "logged in"
}
```

### 3. Messaging Service

#### 3.1 Send Message
**Endpoint:** `/sendMessage`  
**Method:** `POST`  
**Description:** Sends a message from one federation to another.

**Request:**
```json
{
  "from": "string",
  "to": "string",
  "content": "string"
}
```

**Response:**
```json
{
  "messageId": "string",
  "status": "string"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/sendMessage \
  -H "Content-Type: application/json" \
  -d '{
        "from": "user1@fed-1",
        "to": "user2@fed-2",
        "content": "Hello from Federation 1!"
      }'
```

**Example Response:**
```json
{
  "messageId": "msg-1",
  "status": "sent"
}
```

#### 3.2 Get Messages
**Endpoint:** `/getMessages`  
**Method:** `GET`  
**Description:** Retrieves all messages for the federation.

**Response:**
```json
{
  "messages": [
    {
      "from": "string",
      "to": "string",
      "content": "string"
    }
  ]
}
```

**Example Request:**
```bash
curl -X GET http://localhost:8080/getMessages
```

**Example Response:**
```json
{
  "messages": [
    {
      "from": "user1@fed-1",
      "to": "user2@fed-2",
      "content": "Hello from Federation 1!"
    }
  ]
}
```