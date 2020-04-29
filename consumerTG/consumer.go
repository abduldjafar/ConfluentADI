package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type Data struct {
	Markdown string `json:"MARKDOWN"`
}

func main() {
	token := flag.String("token", "", "api key untuk akses telegram")
	server := flag.String("server", "localhost", "server untuk listening topic")
	topic := flag.String("topic", "default", "topic yang digunakan")
	chatID := flag.Int64("chatid", 985052364, "chat id telegram")
	offsets := flag.String("offset", "earliest", "offset untuk costumer")
	flag.Parse()

	var dataTg Data

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": *server,
		"group.id":          *topic + "-group",
		"auto.offset.reset": *offsets,
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{*topic, "^aRegex.*[Tt]opic"}, nil)
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			err := json.Unmarshal(msg.Value, &dataTg)
			if err == nil {
				fmt.Println(dataTg)
				msgT := tgbotapi.NewMessage(*chatID, dataTg.Markdown)
				bot.Send(msgT)
			}

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
