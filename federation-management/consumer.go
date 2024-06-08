package main

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

// FederationEvent represents the structure of the event
type FederationEvent struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

// Kafka variables
var localKafkaBrokers = []string{"localhost:9092"}
var KafkaTopic = "federation-events"

// ConsumeEvents starts consuming events from Kafka
func ConsumeEvents() {
	consumer, err := sarama.NewConsumer(localKafkaBrokers, nil)
	if err != nil {
		log.Fatal("Failed to start Kafka consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Failed to start partition consumer:", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		var event FederationEvent
		if err := json.Unmarshal(message.Value, &event); err != nil {
			log.Println("Failed to unmarshal event:", err)
			continue
		}

		// Handle event (update local state)
		log.Printf("Consumed event: %+v", event)

		switch event.Action {
		case "FederationCreated":
			// Update local state with new federation
			// Implement logic to add federation to db

			return
		}

	}
}

func handleFederationsCreated(data interface{}) {
	// Implement logic
	log.Println("Hnadling FederationCreated event with data:", data)
}
