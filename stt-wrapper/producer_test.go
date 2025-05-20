//go:build integration

// run with `go test -tags=integration ./...`

package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic   = "test-topic"
	broker  = "localhost:9092"
	groupID = "test-group"
)

// Produce sends a message to the Kafka topic
func Produce(t *testing.T, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("sensor"),
		Value: []byte(message),
	})
	if err != nil {
		t.Fatalf("Failed to write message: %v", err)
	}
	t.Log("Message produced")
}

// Consume reads one message from the Kafka topic
func Consume(t *testing.T) string {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}
	t.Logf("Consumed message: %s", msg.Value)
	return string(msg.Value)
}

// TestKafkaEndToEnd verifies that a produced message is consumed correctly
func TestKafkaEndToEnd(t *testing.T) {
	message := fmt.Sprintf("temperature=%.1fC", 20.0+float64(time.Now().UnixNano()%10))
	Produce(t, message)
	received := Consume(t)

	if received != message {
		t.Fatalf("Mismatch: expected %q, got %q", message, received)
	}
	t.Log("Message verified")
}
