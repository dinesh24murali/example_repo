# How to run?

There are 3 docker files in this project each of them use a different docker image. All of these files work, they have the same set of commands to run them.

**docker-compose.yml**
- This file uses the official Kafka image from apache (apache/kafka:3.8.0). With Kafka version 4 they are planning to remove completely and only use `KRaft`
- It does support `KRaft` which is zookeeper's replacement.

**docker-compose-apache-zookeeper.yml**
- This file uses the official Kafka image from apache (apache/kafka:3.8.0) and the official zookeeper image (zookeeper:3.9.2). With Kafka version 4 they are planning to remove completely and only use `KRaft`.

**docker-compose-multi-node.yml**
- This file uses the official Kafka image from apache (apache/kafka:3.8.0) with multiple nodes (1 controller node and 2 broker nodes)
- The consumers and the producers need to use the port exposed by the `broker` nodes. In this example the two brokers are running at ports `9092`, and `39092`. You can use any of these two instants/ports in the `broker` section in `consumer.js` and `producer.js`

**docker-compose-wurstmeister.yml**
- This file uses the kafka image `wurstmeister/kafka` which is 2 years old. I got this from `kafkajs` documentation.
- There is a image for zookeeper that works out of the box with this image called `wurstmeister/zookeeper`. The `wurstmeister/zookeeper` image is 6 years old.

**docker-compose-zookeeper.yml**
- This file uses the kafka image `confluentinc/cp-kafka` which is the updated image. It also has the updated zookeeper counter part image `confluentinc/cp-zookeeper`.
- [confluent](https://www.confluent.io/?session_ref=https://hub.docker.com/&_ga=2.116374657.64705551.1728091888-1214419537.1728091888&_gl=1*11s52fw*_gcl_au*NjQ2MjUwNDAzLjE3MjgwOTE4ODg.*_ga*MTIxNDQxOTUzNy4xNzI4MDkxODg4*_ga_D2D3EGKSGD*MTcyODE0ODUxNS40LjEuMTcyODE0ODU2Ny44LjAuMA..) is the company who invented `Kafka`.

## Steps to run:

**Step 1:** Run the following command to run the docker compose file.
```bash
docker compose -f docker-compose.yml up
```
**Step 2:** Run the consumer
```bash
node consumer.js
```
This might crash at the first time because it takes a few seconds for kafka and zookeeper to get ready. It will work the second time you run the script.

**Step 3:** Run the producer in a new terminal
```bash
node producer.js
```
If you switch to the consumer you should be able to see a event getting logged
```json
{ "value": "Hello KafkaJS junk!" }
```
