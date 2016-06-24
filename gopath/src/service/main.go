package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	var (
		kafkaService = getConfig("KAFKA_SERVICE", "kafka") // The DNS name for input Kafka broker service
		kafkaPort    = getConfig("KAFKA_PORT", "9092")     // Port to connect to input Kafka peers
		topic        = getConfig("TOPIC", "mytopic")       // The topic to consume
		verbose      = getConfig("VERBOSE", "false")       // Set to `true` if you want to turn on sarama logging
	)
	messages := make(chan string)
	saramaProducer := producer(kafkaService, kafkaPort, topic, messages, verbose == "true")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		messages <- scanner.Text()
	}
	_ = saramaProducer.Close()
}

func getConfig(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		result = defaultValue
	}
	log.Printf("%s=%s\n", key, result)
	return result
}
