package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

func main() {

	topic := flag.String("topic", "default", "kafka topic")
	server := flag.String("server", "localhost", "kafka bootstrap server")
	data := flag.String("data", "data.dat", "file that contain datas for produce to kafka")

	flag.Parse()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *server})
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
	datas, err := os.Open(*data)
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(datas)
	for scanner.Scan() {
		for _, word := range []string{scanner.Text()} {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
				Value:          []byte(word),
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 100000)
}
