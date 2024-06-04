## Federations PoC

### How to run

- API Gateway runs on port 8082.

```bash
cd api-gateway
go run main.go
```

- Federation Manager runs on port 8081.

```bash
cd federation/manager
go run main.go
```

### IEEE-2302-2021 :: FHS Member API

As defined in IEEE-2302-2021, FHS Member API Tag 4: Federation. 
FedAdmins can create and manage federations. This collection of calls enables the creation and termination of federations.

#### 1. GET /federations

- Returns a list of all federations
- Method: GET
- Response: A JSON array of Federation objects


#### 2. POST /federations

- Creates a new federation
- Method: POST
- Request Body: A JSON object with the following properties:

```yaml
    name: The name of the new federation
```

- Response: A JSON object representing the created federation, including its ID and name

#### 3. GET /federations/{fed_id}

- Returns details about a specific federation
- Method: GET
- URL Parameter: fed_id (the ID of the desired federation)
- Response: A JSON object representing the requested federation, including its ID and name

#### 4. DELETE /federations/{fed_id}

- Deletes a specific federation
- Method: DELETE
- URL Parameter: fed_id (the ID of the desired federation to delete)
- Response: HTTP 204 No Content if the deletion is successful

#### Federation instance

The Federation object has the following properties:

- id: Unique ID of the federation
- name: Name of the federation
- service_catalogue: An array of Service objects

### IEEE-2302-2021 :: FHS Operator API

#### 1. POST /FHSOperator/NewFedAdmin

- Creates a new Fed Admin
- Method: POST
- Request: a JSON object with the following properties:

```yaml
    name: The name of the new Fed Admin (required)
    email: The email address of the new Fed Admin (optional)
    description: A brief description of the new Fed Admin (optional)
    enabled: Whether the new Fed Admin is enabled or not (required)
```

- Response: a JSON object representing the created Fed Admin, including:

```yaml
    member_id: The unique ID assigned to the new Fed Admin
    member_name: The name of the new Fed Admin
    email: The email address of the new Fed Admin (if provided)
    description: A brief description of the new Fed Admin (if provided)
    enabled: Whether the new Fed Admin is enabled or not
```

### Notes

- All endpoints use JSON as the response format.
- Error responses will be sent with HTTP status codes indicating the error type (e.g., 400 Bad Request, 404 Not Found, etc.).