# Kafka

## Research:

Following are some of the client libraries for NodeJS

1. [kafkajs](https://www.npmjs.com/package/kafkajs)
2. [kafka-node](https://www.npmjs.com/package/kafka-node)
3. [node-rdkafka](https://github.com/Blizzard/node-rdkafka)

For `TYS project` we were using `node-rdkafka`. We had a few problems with missed kafka events.

## [How does Kafka work in a nutshell?](https://kafka.apache.org/documentation/#intro_nutshell)
### Servers: 
Kafka is run as a cluster of one or more servers that can span multiple data-centers or cloud regions. Some of these servers form the storage layer, called the `brokers`.

### Clients:
They allow you to write distributed applications and microservices that read, write, and process streams of events in parallel, at scale, and in a fault-tolerant manner even in the case of network problems or machine failures.

Producers are those client applications that publish (write) events to Kafka, and consumers are those that subscribe to (read and process) these events.

# How to run?

```bash
export HOST_IP=$(ifconfig | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3}" | grep -v 127.0.0.1 | awk '{ print $2 }' | cut -f2 -d: | head -n1)
```
```bash
docker-compose up
```