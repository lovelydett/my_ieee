// Package components provides kafka consumer and producer
package components

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer is a wrapper around the kafka consumer
type KafkaConsumer struct {
	// consumer is the underlying kafka consumer
	consumer *kafka.Consumer
}

// NewKafkaConsumer creates a new kafka consumer
func NewKafkaConsumer(brokers, groupID, topic string) (*KafkaConsumer, error) {
	// create a new kafka consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	// subscribe to the topic
	c.Subscribe(topic, nil)
	return &KafkaConsumer{consumer: c}, nil
}
