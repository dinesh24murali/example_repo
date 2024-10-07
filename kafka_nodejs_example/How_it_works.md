# How this integration is working:

The official Kafka image is from apache. There are many other companies that have open sourced their own images. One of the popular one is `confluentinc/cp-zookeeper` and `wurstmeister/kafka`. `confluentinc/cp-zookeeper` is being actively updated at the time of writing this document (7-10-2024).

This project has multiple docker compose files for all these images. The environment variables for these kafka images are different, meaning an environment variable used for one image might not work for another image.

Since kafka has decided to [deprecate zookeeper](https://kafka.apache.org/documentation/#zk_depr) I decided to use the [official kafka image](https://hub.docker.com/r/apache/kafka) form apache that uses `Kraft`. They are planning to remove zookeeper all together in version 4. This is implemented in the following two files:
1. docker-compose.yml
2. docker-compose-multi-node.yml

There are 2 types of nodes in kafka they are `broker` and `controller`. Kafka can be deployed with multiple `broker` nodes and `controller` nodes. When we initialize the cluster, all the available kafka `controller` nodes in the network will do a vote to elect the `controller node`.  Even if you have one kafka server in you cluster, it will still do the vote to select the `controller node`. This takes a few seconds to decide the `controller node`. The election process will get triggered after the server is accessed for the first time after initialization. So it is best to initialize the consumer first, and not trigger the producer immediately after the server starts.

We can manually assign a role to a node. In this case the single node acts as both the broker and the controller.

This [stackoverflow thread](https://stackoverflow.com/questions/42998859/kafka-server-configuration-listeners-vs-advertised-listeners) explains the difference between `KAFKA_LISTENER_SECURITY_PROTOCOL_MAP` and `KAFKA_LISTENERS`.

### Single node:
```yml
version: "3.9"
services:
  broker:
    image: apache/kafka:3.8.0
    hostname: broker
    container_name: broker
    ports:
      - '9092:9092'
    environment:
      KAFKA_NODE_ID: 1
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT_HOST://localhost:9092,PLAINTEXT://broker:19092'
      KAFKA_PROCESS_ROLES: 'broker,controller'
      KAFKA_CONTROLLER_QUORUM_VOTERS: '1@broker:29093'
      KAFKA_LISTENERS: 'CONTROLLER://:29093,PLAINTEXT_HOST://:9092,PLAINTEXT://:19092'
      KAFKA_INTER_BROKER_LISTENER_NAME: 'PLAINTEXT'
      KAFKA_CONTROLLER_LISTENER_NAMES: 'CONTROLLER'
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
      KAFKA_LOG_DIRS: '/tmp/kraft-combined-logs'
```
### Multi node:
This file has two broker nodes and one controller node.
```yml
version: "3.9"
services:
  controller-1:
    image: apache/kafka:latest
    container_name: controller-1
    environment:
      KAFKA_NODE_ID: 1
      KAFKA_PROCESS_ROLES: controller
      KAFKA_LISTENERS: CONTROLSingleLER://:9093
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 2
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_LOG_DIRS: '/tmp/kraft-combined-logs'
  broker-1:
    image: apache/kafka:latest
    container_name: broker-1
    ports:
      - 9092:9092
    environment:
      KAFKA_NODE_ID: 2
      KAFKA_PROCESS_ROLES: broker
      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-1:19092,PLAINTEXT_HOST://localhost:9092'
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 2
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_LOG_DIRS: '/tmp/kraft-combined-logs'
    depends_on:
      - controller-1
  broker-2:
    image: apache/kafka:latest
    container_name: broker-2
    ports:
      - 39092:9092
    environment:
      KAFKA_NODE_ID: 3
      KAFKA_PROCESS_ROLES: broker
      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-2:19092,PLAINTEXT_HOST://localhost:39092'
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 2
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_LOG_DIRS: '/tmp/kraft-combined-logs'
    depends_on:
      - controller-1
```

- **KAFKA_NODE_ID**: This needs to be unique for each kafka node in the network, since there is only a single service in this file it doesn't matter here.
- **KAFKA_LISTENER_SECURITY_PROTOCOL_MAP**: defines key/value pairs for the security protocol to use per listener name. Here `PLAINTEXT` means `http`.
- **KAFKA_ADVERTISED_LISTENERS**: this specifies the port through which the listeners (both consumers and producers) will use.
- **KAFKA_CONTROLLER_QUORUM_VOTERS**: this specifier the list of controller nodes that are present in the system. Kafka will use this to identify and elect the primary controller node.
- **KAFKA_LISTENERS**: this specifies the port the broker node listens on. This [stackoverflow thread](https://stackoverflow.com/questions/42998859/kafka-server-configuration-listeners-vs-advertised-listeners) explains the difference between `KAFKA_LISTENER_SECURITY_PROTOCOL_MAP` and `KAFKA_LISTENERS`. I believe this data gets stored in cache/DB.
- **KAFKA_INTER_BROKER_LISTENER_NAME** & **KAFKA_CONTROLLER_LISTENER_NAMES**: Kafka brokers communicate between themselves, usually on the internal network (e.g., Docker network, AWS VPC, etc.). To define which listener to use, specify KAFKA_INTER_BROKER_LISTENER_NAME(inter.broker.listener.name). The host/IP used must be accessible from the broker machine to others.
- **CLUSTER_ID**: This is the unique id the represents a cluster. This needs to be the same in all nodes when we are running multiple instances. This gets generated automatically. I have hardcoded it here so that when we restart the services the old messages is still accessible.
- **KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR**: This represent the number of broker/replica nodes we have in the system. The default value is 3, so you need to explicitly set this value based on your setup.
- **KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS**: reduce the initial delay you see when launching a KSQL query, this makes the whole thing feel much faster. (the default value, iirc, is 3000) For some reason the `docker-compose-multi-node.yml` worked with wait time as zero.
- **KAFKA_LOG_DIRS**: The location of the config and log files.
- **KAFKA_ZOOKEEPER_CONNECT**: This is where we specify the URL to connect to zookeeper.

# Reference links

https://hub.docker.com/r/apache/kafka
Multi node example (not tested), it was taken from the above link.
```yml
services:
  controller-1:
    image: apache/kafka:latest
    container_name: controller-1
    environment:
      KAFKA_NODE_ID: 1
      KAFKA_PROCESS_ROLES: controller
      KAFKA_LISTENERS: CONTROLLER://:9093
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093,3@controller-3:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0

  controller-2:
    image: apache/kafka:latest
    container_name: controller-2
    environment:
      KAFKA_NODE_ID: 2
      KAFKA_PROCESS_ROLES: controller
      KAFKA_LISTENERS: CONTROLLER://:9093
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093,3@controller-3:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0

  controller-3:
    image: apache/kafka:latest
    container_name: controller-3
    environment:
      KAFKA_NODE_ID: 3
      KAFKA_PROCESS_ROLES: controller
      KAFKA_LISTENERS: CONTROLLER://:9093
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093,3@controller-3:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0

  broker-1:
    image: apache/kafka:latest
    container_name: broker-1
    ports:
      - 29092:9092
    environment:
      KAFKA_NODE_ID: 4
      KAFKA_PROCESS_ROLES: broker
      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-1:19092,PLAINTEXT_HOST://localhost:29092'
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093,3@controller-3:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
    depends_on:
      - controller-1
      - controller-2
      - controller-3

  broker-2:
    image: apache/kafka:latest
    container_name: broker-2
    ports:
      - 39092:9092
    environment:
      KAFKA_NODE_ID: 5
      KAFKA_PROCESS_ROLES: broker
      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-2:19092,PLAINTEXT_HOST://localhost:39092'
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093,3@controller-3:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
    depends_on:
      - controller-1
      - controller-2
      - controller-3

  broker-3:
    image: apache/kafka:latest
    container_name: broker-3
    ports:
      - 49092:9092
    environment:
      KAFKA_NODE_ID: 6
      KAFKA_PROCESS_ROLES: broker
      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-3:19092,PLAINTEXT_HOST://localhost:49092'
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093,3@controller-3:9093
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
    depends_on:
      - controller-1
      - controller-2
      - controller-3
```