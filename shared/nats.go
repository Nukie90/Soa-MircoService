package shared

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// ConnectNATS establishes a connection to the NATS server and returns the connection.
func ConnectNATS() (*nats.Conn, error) {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL // Use the default URL as fallback
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to NATS")
	return nc, nil
}

func MarshalToJSON(data interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling to JSON: %v", err)
		return nil
	}
	return jsonData
}
