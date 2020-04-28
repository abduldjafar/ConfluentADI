package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	email := flag.String("email", "", "email for get data")
	passwd := flag.String("password", "", "password for get data")
	topic := flag.String("topic", "asoex", "topic kafka untuk menyimpan data")
	datakind := flag.String("datakind", "cif", "jenis data yang akan di produce")
	key := flag.String("openkey", "", "key yang digunakan untuk mengambil data location")
	listtempat := flag.String("listTempat", "stasiun", "file list tempat untuk opencage")
	host := flag.String("host", "localhost", "host for kafka")
	flag.Parse()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *host})
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
		// Produce messages to topic (asynchronously)
		for scanner.Scan() {
			line := scanner.Text()
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
				Value:          []byte(line),
			}, nil)
			p.Flush(15 * 100000)

		}

	} else if *datakind == "opencage" && *listtempat != "" && *key != "" {
		data, err := os.Open(*listtempat)
		if err != nil {
			log.Println(err.Error())
		}
		scanner := bufio.NewScanner(data)
		for scanner.Scan() {
			tempat := "stasiun%20" + strings.Trim(scanner.Text(), " ")
			url := "https://api.opencagedata.com/geocode/v1/json?key=" + *key + "&q=" + tempat

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Println(err.Error())
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err.Error())
			}
			defer resp.Body.Close()
			log.Println("response " + resp.Status)
			body, _ := ioutil.ReadAll(resp.Body)
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
				Value:          []byte(body),
			}, nil)
			p.Flush(15 * 100000)
		}
	}
}
