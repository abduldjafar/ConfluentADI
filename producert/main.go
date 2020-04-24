package main

import (
	"bufio"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "xxxx6"
	data, err := os.Open("data")
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		for _, word := range []string{scanner.Text()} {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(word),
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 100000)
}
