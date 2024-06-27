package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	multicastAddr = "239.0.1.100:12345"
)

type Instance struct {
	InstanceID       *int   `json:"instance_id"`
	InstanceIP       string `json:"instance_ip"`
	InstancePort     *int   `json:"instance_port"`
	InstanceMetadata string `json:"instance_metadata"`
}

type FCPMessage struct {
	Type     string   `json:"type"`
	Instance Instance `json:"instance"`
}

var (
	currentInstance Instance
	catalogue       = make(map[int]Instance)
	catalogueMutex  sync.RWMutex
)

func main() {
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

func handleRegister(instance Instance) {
	// First we check if the instance is ourselves
	if *instance.InstanceID == *currentInstance.InstanceID {
		log.Printf("Received REGISTER message from self (ID %d). Ignoring.", *instance.InstanceID)
		return
	}

	isNewInstance := false
	catalogueMutex.Lock()
	if _, exists := catalogue[*instance.InstanceID]; !exists {
		isNewInstance = true
		catalogue[*instance.InstanceID] = instance
		log.Printf("Registered new instance: ID %d, IP %s, Port %d", *instance.InstanceID, instance.InstanceIP, *instance.InstancePort)
	}
	catalogueMutex.Unlock()

	if isNewInstance {
		// Send current instance information to the new instance
		sendRegisterResponse(instance)
	}
}

func sendRegisterResponse(newInstance Instance) {
	// Check if the new instance is ourselves
	if *newInstance.InstanceID == *currentInstance.InstanceID {
		log.Printf("Attempted to send REGISTER_RESPONSE to ourselves, ignoring")
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
	catalogueMutex.RLock()         // Acquire shared lock for reading
	defer catalogueMutex.RUnlock() // Ensure the lock is released when the function returns

	fmt.Println("Current Instance Catalogue:")
	for id, instance := range catalogue {
		fmt.Printf("ID: %d, IP: %s, Port: %d, Metadata: %s\n", id, instance.InstanceIP, *instance.InstancePort, instance.InstanceMetadata)
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
