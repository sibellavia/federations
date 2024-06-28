package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Define the multicast address for communication
const (
	multicastAddr = "239.0.1.100:12345"
)

// Instance represents a Node instance
type Instance struct {
	InstanceID       *int   `json:"instance_id"`
	InstanceIP       string `json:"instance_ip"`
	InstancePort     *int   `json:"instance_port"`
	InstanceMetadata string `json:"instance_metadata"`
}

// FCPMessage represents a message in the FCP protocol.
type FCPMessage struct {
	Type     string   `json:"type"`
	Instance Instance `json:"instance"`
}

var db *sql.DB

// currentInstance represents the currently active instance/node
var currentInstance Instance

func main() {
	// main initializes the currentInstance struct with values from environment variables,
	// and starts goroutines to listen for multicast messages and handle user input.
	// The main goroutine is kept running indefinitely using a select statement.

	initDB()
	defer db.Close()

	id, _ := strconv.Atoi(os.Getenv("INSTANCE_ID"))
	port, _ := strconv.Atoi(os.Getenv("INSTANCE_PORT"))
	currentInstance = Instance{
		InstanceID:       &id,
		InstanceIP:       os.Getenv("INSTANCE_IP"),
		InstancePort:     &port,
		InstanceMetadata: os.Getenv("INSTANCE_METADATA"),
	}

	if currentInstance.InstanceID == nil || currentInstance.InstanceIP == "" || currentInstance.InstancePort == nil {
		log.Fatal("Instance ID, IP and Port are required")
	}

	go listenMulticast()
	go handleUserInput()

	select {} // keep the main goroutine running
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./fcp.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Create the instances table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS instances (
		instance_id INTEGER PRIMARY KEY,
		instance_ip TEXT NOT NULL,
		instance_port INTEGER NOT NULL,
		instance_metadata TEXT
	)`)
	if err != nil {
		log.Fatalf("Error creating instances table: %v", err)
	}
}

// listenMulticast listens for multicast messages on the specified address.
// It resolves the UDP address, creates a multicast UDP connection, and starts
// a separate goroutine to listen for unicast responses. It reads multicast
// messages, unmarshals them into FCPMessage struct, and handles REGISTER
// messages by calling the handleRegister function.
func listenMulticast() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		log.Fatalf("Error resolving multicast address: %v", err)
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error listening multicast: %v", err)
	}
	defer conn.Close()

	// Start a separate goroutine to listen for unicast responses
	go listenUnicast()

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading multicast: %v", err)
			continue
		}

		var message FCPMessage
		if err := json.Unmarshal(buffer[:n], &message); err != nil {
			log.Printf("Error unmarshalling multicast message: %v", err)
			continue
		}

		if message.Type == "REGISTER" {
			handleRegister(message.Instance)
		}
	}
}

// listenUnicast listens for UDP unicast messages on the specified port.
// It resolves the UDP address, creates a UDP connection, and continuously reads incoming messages.
// If a message is successfully read and unmarshalled, it checks the message type and handles it accordingly.
// This function is designed to be run as a goroutine.
func listenUnicast() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", *currentInstance.InstancePort))
	if err != nil {
		log.Fatalf("Error resolving unicast address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Error listening unicast: %v", err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading unicast: %v", err)
			continue
		}

		var message FCPMessage
		if err := json.Unmarshal(buffer[:n], &message); err != nil {
			log.Printf("Error unmarshalling unicast message: %v", err)
			continue
		}

		if message.Type == "REGISTER_RESPONSE" {
			handleRegister(message.Instance)
		}
	}
}

// sendRegister sends a REGISTER message to a multicast address.
// It resolves the UDP address, dials the multicast address, and sends the message.
// The message contains the type "REGISTER" and the current instance.
// If there is an error resolving the multicast address, dialing it,
// marshaling the JSON, or sending the multicast message, an error is logged.
// Otherwise, a success message is logged.
func sendRegister() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		log.Printf("Error resolving multicast address: %v", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Printf("Error dialing multicast address: %v", err)
		return
	}
	defer conn.Close()

	message := FCPMessage{
		Type:     "REGISTER",
		Instance: currentInstance,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		log.Printf("Error sending multicast message: %v", err)
	} else {
		log.Printf("Sent REGISTER message for our instance (ID %d)", *currentInstance.InstanceID)
	}
}

// handleRegister handles the REGISTER message received from an instance.
// If the instance is new, it adds it to the catalogue and sends the current
// instance information to the new instance.
func handleRegister(instance Instance) {
	if *instance.InstanceID == *currentInstance.InstanceID {
		return
	}

	exists, err := instanceExists(*instance.InstanceID)
	if err != nil {
		log.Printf("Error checking if instance exists: %v", err)
		return
	}

	if exists {
		log.Printf("Instance ID %d already exists in the catalogue", *instance.InstanceID)
		return
	}

	stmt, err := db.Prepare(`
        INSERT OR REPLACE INTO instances (instance_id, instance_ip, instance_port, instance_metadata)
        VALUES (?, ?, ?, ?)
    `)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(*instance.InstanceID, instance.InstanceIP, *instance.InstancePort, instance.InstanceMetadata)
	if err != nil {
		log.Printf("Error inserting instance: %v", err)
		return
	}

	log.Printf("Registered instance: ID %d, IP %s, Port %d", *instance.InstanceID, instance.InstanceIP, *instance.InstancePort)

	// Send current instance information to the new instance
	sendRegisterResponse(instance)
}

// sendRegisterResponse sends a REGISTER_RESPONSE message to a new instance.
// It dials the new instance using UDP and sends the REGISTER_RESPONSE message.
// The function takes a new instance as a parameter.
// It returns an error if there is any issue with dialing or sending the message.
func sendRegisterResponse(newInstance Instance) {
	// Check if the new instance is ourselves
	if *newInstance.InstanceID == *currentInstance.InstanceID {
		return
	}

	exists, err := instanceExists(*currentInstance.InstanceID)
	if err != nil {
		log.Printf("Error checking if instance exists: %v", err)
		return
	}

	if exists {
		log.Printf("Instance ID %d already exists in the catalogue", *currentInstance.InstanceID)
		return
	}

	addr := fmt.Sprintf("%s:%d", newInstance.InstanceIP, *newInstance.InstancePort)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Printf("Error dialing new instance: %v", err)
		return
	}
	defer conn.Close()

	message := FCPMessage{
		Type:     "REGISTER_RESPONSE",
		Instance: currentInstance,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		log.Printf("Error sending register response: %v", err)
	} else {
		log.Printf("Sent REGISTER_RESPONSE to instance ID %d", *newInstance.InstanceID)
	}
}

func printCatalogue() {
	rows, err := db.Query("SELECT instance_id, instance_ip, instance_port, instance_metadata FROM instances")
	if err != nil {
		log.Printf("Error querying instances: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("Current Instance Catalogue:")
	for rows.Next() {
		var id, port int
		var ip, metadata string
		err := rows.Scan(&id, &ip, &port, &metadata)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		fmt.Printf("ID: %d, IP: %s, Port: %d, Metadata: %s\n", id, ip, port, metadata)
	}
}

func handleUserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter 'print' to view the catalogue, 'register' to send a REGISTER message, or 'quit' to exit:")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "print":
			printCatalogue()
		case "register":
			sendRegister()
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}

func instanceExists(id int) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM instances WHERE instance_id = ?)", id).Scan(&exists)
	return exists, err
}
