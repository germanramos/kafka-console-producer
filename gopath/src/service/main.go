package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/Shopify/sarama.v1"
)

var (
	brokers   = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
	verbose   = flag.Bool("verbose", false, "Turn on Sarama logging")
	certFile  = flag.String("certificate", "", "The optional certificate file for client authentication")
	keyFile   = flag.String("key", "", "The optional key file for client authentication")
	caFile    = flag.String("ca", "", "The optional certificate authority file for TLS client authentication")
	verifySsl = flag.Bool("verify", false, "Optional verify ssl certificates chain")
	topic     = flag.String("topic", "kafka-producer", "Topic to produce to")
)

func main() {
	flag.Parse()

	if *verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	//producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	var producer sarama.SyncProducer
	var err error
	for producer == nil {
		producer, err = sarama.NewSyncProducer(brokerList, config)
		if err != nil {
			log.Println("Failed to start Sarama producer:", err)
		}
		time.Sleep(time.Second * 3)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, _, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(scanner.Text()),
		})
		if err != nil {
			log.Println("Failed to produce:", err)
		}
	}
}
