### kafka-console-producer

Golang kafka-console-producer using sarama.

It produces messages reading from stdin

## Usage

```
./main
```

## Configuration

Use environment variables

- KAFKA_SERVICE, "kafka", The DNS name for input Kafka broker service
- KAFKA_PORT, "9092", Port to connect to input Kafka peers
- TOPIC, "kafka-console-producer", The topic to consume
- VERBOSE, "false, Set to `true` if you want verbose output
