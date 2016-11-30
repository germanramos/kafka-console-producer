package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)
	log.Print("kafka-console-producer v0.7.0")
	var (
		kafkaService  = getConfig("KAFKA_SERVICE", "kafka") // The DNS name for input Kafka broker service
		kafkaPort     = getConfig("KAFKA_PORT", "9092")     // Port to connect to input Kafka peers
		topic         = getConfig("TOPIC", "mytopic")       // The topic to produce to
		partition     = getConfig("PARTITION", "-1")        // The partition to produce to. All if not defined.
		verbose       = getConfig("VERBOSE", "false")       // Set to `true` if you want to turn on sarama logging
		finishTimeout = getConfig("FINISH_TIMEOUT", "1")    // Number of seconds to wait for exit after end of line is received
		key           = getConfig("KEY", "")                // The key of produced messages. Default (empty) will produce in every topic
	)
	messages := make(chan string)
	partitionInt, _ := strconv.Atoi(partition)
	saramaProducer := producer(kafkaService, kafkaPort, topic, int32(partitionInt), messages, key, verbose == "true")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		event := scanner.Text()
		if event != "" {
			messages <- scanner.Text()
		} else {
			log.Print("Discarting empty line")
		}
	}
	finishTimeoutInt, _ := strconv.Atoi(finishTimeout)
	time.Sleep(time.Duration(finishTimeoutInt) * time.Second)
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
