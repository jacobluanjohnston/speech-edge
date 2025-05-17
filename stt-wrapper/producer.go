package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func sendTestMessage() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := kafka.Message{
		Key:   []byte("transcript"),
		Value: []byte(`{"text": "hello world", "confidence": 0.91}`),
	}

	fmt.Println("Sending message to Kafka...")
	err := w.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatalf("could not write message: %v", err)
	}

	fmt.Println("Message sent!")
	time.Sleep(1 * time.Second)
}
