package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net/http"
)

func main() {
	email := flag.String("email", "hengkiriang@alldataint.com", "email for get data")
	passwd := flag.String("password", "", "password for get data")
	topic := flag.String("topic", "asoex", "topic kafka untuk menyimpan data")
	datakind := flag.String("datakind", "cif", "jenis data yang akan di produce")
	flag.Parse()

	var scanner *bufio.Scanner
	if *datakind == "cif" {
		req, err := http.NewRequest("GET", "https://datafeeds.networkrail.co.uk/ntrod/CifFileAuthenticate?type=CIF_EA_TOC_FULL_DAILY&day=toc-full", nil)
		if err != nil {
			// handle err
		}
		req.SetBasicAuth(*email, *passwd)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
		}
		defer resp.Body.Close()

		gReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(gReader)

	} else if *datakind == "opencage" {

	}

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
	for scanner.Scan() {
		line := scanner.Text()
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
			Value:          []byte(line),
		}, nil)

	}

	p.Flush(15 * 100000)
}
