package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"gopkg.in/Shopify/sarama.v1"
)

func producer(kafkaService string,
	kafkaPort string,
	topic string,
	verbose bool) {

	var (
		err        error
		brokerList []string
		producer   sarama.AsyncProducer
	)

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	if kafkaService == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Get Kafka peers
	for err != nil || brokerList == nil {
		brokerList, err = net.LookupHost(kafkaService)
		if err != nil {
			log.Printf("Failed to resolve %s: %s\n", kafkaService, err)
			time.Sleep(time.Second * 3)
		}
	}
	for i, e := range brokerList {
		brokerList[i] = e + ":" + kafkaPort
	}
	log.Println("Kafka Broker List:", brokerList)

	//producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5 // Retry up to 5 times to produce the message

	for producer == nil {
		producer, err = sarama.NewAsyncProducer(brokerList, config)
		if err != nil {
			log.Println("Failed to start Sarama producer:", err)
			time.Sleep(time.Second * 3)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(scanner.Text()),
		}
		producer.Input() <- msg
	}
}
