package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	var (
		kafkaService = getConfig("KAFKA_SERVICE", "kafka")          // The DNS name for input Kafka broker service
		kafkaPort    = getConfig("KAFKA_PORT", "9092")              // Port to connect to input Kafka peers
		topic        = getConfig("TOPIC", "kafka-console-producer") // The topic to consume
		verbose      = getConfig("VERBOSE", "false")                // Set to `true` if you want to turn on sarama logging
	)
	producer(kafkaService, kafkaPort, topic, verbose == "true")
}

func getConfig(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		result = defaultValue
	}
	log.Printf("%s=%s\n", key, result)
	return result
}
