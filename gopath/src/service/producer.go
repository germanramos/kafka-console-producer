package main

import (
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
	messages chan string,
	key string,
	verbose bool) sarama.AsyncProducer {

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
	config.Producer.RequiredAcks = sarama.WaitForLocal  // The level of acknowledgement reliability needed from the broker
	config.Producer.Timeout = 1000 * time.Millisecond   // The maximum duration the broker will wait the receipt of the number of RequiredAcks
	config.Producer.Return.Errors = true                // If enabled, messages that failed to deliver will be returned on the Errors channel
	config.Producer.Retry.Max = 5                       // Retry up to 5 times to produce the message
	config.Metadata.Retry.Max = 5000                    // The total number of times to retry a metadata request when the cluster is in the middle of a leader election
	config.Metadata.Retry.Backoff = 1 * time.Second     // How long to wait for leader election to occur before retrying (default 250ms)
	config.Metadata.RefreshFrequency = 30 * time.Second // How frequently to refresh the cluster metadata in the background. Defaults to 10 minutes.
	for producer == nil {
		producer, err = sarama.NewAsyncProducer(brokerList, config)
		if err != nil {
			log.Println("Failed to start Sarama producer:", err)
			time.Sleep(time.Second * 3)
		}
	}

	// log failed messages
	go func() {
		for {
			var msg *sarama.ProducerError
			msg = <-producer.Errors()
			log.Printf("Error producing message: %s", msg.Msg.Value)
		}
	}()

	// Read stdin for ever and publish messages
	var producerKey sarama.Encoder
	if key == "" {
		producerKey = nil
	} else {
		producerKey = sarama.StringEncoder(key)
	}

	go func() {
		for {
			msg := <-messages
			producerMessage := &sarama.ProducerMessage{
				Topic: topic,
				Key:   producerKey,
				Value: sarama.StringEncoder(msg),
			}
			producer.Input() <- producerMessage
		}
	}()

	return producer

}
