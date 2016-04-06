### kafka-console-producer

Golang kafka-console-producer using sarama.

It produces messages reading from stdin

## Usage

```
Usage of ./main:
  -brokers string
    	The Kafka brokers to connect to, as a comma separated list
  -ca string
    	The optional certificate authority file for TLS client authentication
  -certificate string
    	The optional certificate file for client authentication
  -key string
    	The optional key file for client authentication
  -topic string
    	Topic to produce to (default "kafka-producer")
  -verbose
    	Turn on Sarama logging
  -verify
    	Optional verify ssl certificates chain
```

If no brokers specified it will try to read KAFKA_PEERS env variable. Example:

KAFKA_PEERS=kafka_broker_1.broker.kafka.skydns.local:9092
