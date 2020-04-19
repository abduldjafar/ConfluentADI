package main

import (
	"ConfluentADI/config"
	"ConfluentADI/model"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type user struct {
	Word   string
	Jumlah int
}

func main() {

	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	db, err := gorm.Open("postgres", "host="+baseConfig.Postgres.Url+" port="+baseConfig.Postgres.Port+""+
		" user="+baseConfig.Postgres.User+" dbname="+baseConfig.Postgres.Db+" password="+baseConfig.Postgres.Password+
		" sslmode=disable")

	tbwordcount := model.DBMigrationAccount(db, &model.WordCount{})

	wordcount := model.WordCount{}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myTopic-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"xxxx3-group-table", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var obj []user
			if err := json.Unmarshal(msg.Value, &obj); err != nil {
				log.Println(err.Error())
			}

			for _, dict := range obj {
				wordcount.Word = dict.Word
				wordcount.Jumlah = dict.Jumlah

				if err := tbwordcount.Save(&wordcount).Error; err != nil {
					log.Println(err.Error())
				}

			}

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
