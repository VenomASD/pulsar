package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {
	//client code
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	//consumer code
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Topic() string
	// ProducerName() string
	// Properties() map[string]string
	// Payload() []byte

	// PublishTime() time.Time
	// EventTime() time.Time
	// Key() string
	// OrderingKey() string
	// RedeliveryCount() uint32
	// IsReplicated() bool
	// GetReplicatedFrom() string
	// GetSchemaValue(v interface{}) error
	// GetEncryptionContext() *EncryptionContext
	fmt.Printf("Received message msgId: %#v -- content: '%s'\n Topic: %s %s",
		msg.ID(), string(msg.Payload()),msg.Topic(),msg.Topic())

	consumer.Ack(msg)

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
	consumer.Close()
	client.Close()
}