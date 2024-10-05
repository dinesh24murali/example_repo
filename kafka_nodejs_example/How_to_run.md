# How to run?

There are 3 docker files in this project each of them use a different docker image. All of these files work, they have the same set of commands to run them.

**docker-compose-apache-KRaft.yml**
- This file uses the official Kafka image from apache (apache/kafka:3.8.0).
- This image either doesn't support Kafka zookeeper setup out of the box or I can't find the instructions
- It does support `KRaft` which is zookeeper's replacement

**docker-compose-wurstmeister.yml**
- This file uses the kafka image `wurstmeister/kafka` which is 2 years old. I got this from `kafkajs` documentation.
- There is a image for zookeeper that works out of the box with this image called `wurstmeister/zookeeper`. The `wurstmeister/zookeeper` image is 6 years old.

**docker-compose.yml**
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
```bash
{ value: 'Hello KafkaJS junk!' }
```
