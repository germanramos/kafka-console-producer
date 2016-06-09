### kafka-console-producer

kafka-console-producer implemented in golang and using [sarama](https://github.com/Shopify/sarama) driver.

## Features

- Works with Apache Kafka >= 0.8
- Statically compiled. No dependencies. It run on every linux distribution.
- Just one binary file of ~ 6 Mb
- It produces messages reading from stdin
- It prints state messages in stderr
- Very easy to configure trough environment variables
- Auto discover kafka peers from DNS name
- Waits for kafka to be ready
- Auto reconnect

## Usage

```
./kafka-console-producer
```

## Configuration

Use environment variables

- KAFKA_SERVICE, "kafka", The DNS name for input Kafka broker service
- KAFKA_PORT, "9092", Port to connect to input Kafka peers
- TOPIC, "kafka-console-producer", The topic to consume
- VERBOSE, "false, Set to `true` if you want verbose output

## Download

https://github.com/germanramos/kafka-console-producer/releases/download/v0.3.0/kafka-console-producer

## Run Example

```
KAFKA_SERVICE=192.168.1.45 TOPIC=foo ./kafka-console-consumer
```

## Related work

https://github.com/germanramos/kafka-console-consumer

## License

MIT - German Ramos
