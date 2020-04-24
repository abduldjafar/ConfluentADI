package main

import (
	"ConfluentADI/config"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type user struct {
	Word   string
	Jumlah int
}

func main() {

	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myTopic-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"testopencage", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			fmt.Println(string(msg.Value))

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
