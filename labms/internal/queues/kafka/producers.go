package kfka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	topic    string
	producer *kafka.Producer
}

type CustomMessage struct {
	Message string
	Data    interface{}
	Status  string
}

func NewKafkaProducer(t, broker string) *KafkaProducer {
	np, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		log.Println("ERROR in creating producer", err)
	}
	return &KafkaProducer{
		topic:    t,
		producer: np,
	}
}

func (k *KafkaProducer) ProduceMessage(msg CustomMessage) error {
	ubyte, e := json.Marshal(msg)
	if e != nil {
		fmt.Println("ERROR : UserSQLRepo Marshalling error  ", e)
	}

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.topic,
			Partition: kafka.PartitionAny,
		},
		Value: ubyte,
	}, nil)
	if err != nil {
		log.Println("ERROR : ProduceMessage in topic - ", k.topic, "-- error - ", err)
		return err
	}
	k.producer.Flush(1000) // waits 1 second for delivery

	return nil
}
