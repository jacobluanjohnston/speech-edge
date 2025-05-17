package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	sendTestMessage()

	// Start consuming
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		GroupID:  "stt-wrapper-group",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	defer r.Close()

	fmt.Println("Listening for messages on topic: test-topic")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("could not read message: %v", err)
		}
		fmt.Printf("Received message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
