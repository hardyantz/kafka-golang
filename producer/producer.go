package producer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Email   string `json:"email"`
}

func SendMessage(m Message) error {
	topic := "my-message"

	conn, err := kafka.NewProducer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		return err
	}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = conn.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(m.Email),
		Value: []byte(b),
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
