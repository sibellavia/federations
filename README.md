## Federations PoC

### How to run

- API Gateway runs on port 8082.

```bash
cd api-gateway
go run main.go
```

- Federation Manager runs on port 8081.

```bash
cd federation-management
go run main.go
```

- Federation Member Management runs on port 8083.

```bash
cd federation-member-management
go run main.go
```

- Kafka

1. Follow the instructions on the [Kafka quickstart guide](https://kafka.apache.org/quickstart) to download and install Kafka on each VM where the PoC is running.

2. Start Kafka Server:

```bash
bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties
```

3. Create a topic for federation events on each Kafka instance:

```bash
bin/kafka-topics.sh --create --topic federation-events --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
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

### Relationships

#### fed_admins to federations: One-to-Many relationship.

Each Federation Administrator (fed_admins.member_id) can create multiple Federations (federations.member_id).

#### federations to services: One-to-Many relationship.

Each Federation (federations.fed_id) can have multiple Services (services.fed_id).

### Note per me stesso (per ricordarmi)

Per avere un'architettura completamente federata, dove ciascuna istanza dell'OVP sia autonoma, indipendente e self-contained, e comunque in grado di poter comunicare e condividere dati con altre istanze, dovrei implementare un sistema distribuito di brokering.

Quindi l'idea di implementazione è:
1. Deploy di istanze kafka indipendenti (ogni OVP avrà la sua istanza con il proprio cluster di Kafka -> sicuramente dovrò implementare una replicazione cross-cluster di Kafka per sincronizzare gli eventi delle federazioni fra gli OVP?)
2. Implementazione pub/sub (ogni OVP pubblicherà eventi relativi alle federazioni sul proprio cluster di Kafka, e ogni OVP consumerà eventi per aggiornare il proprio local state)
3. Implementazione cross-cluster replication (kafka mirrormaker per replicare eventi fra i vari cluster di kafka? eventi che andrebbero replicati sono tipo federation creation, updates, joins...)