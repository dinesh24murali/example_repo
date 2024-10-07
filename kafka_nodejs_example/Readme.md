# Kafka

## Research:

Following are some of the client libraries for NodeJS

1. [kafkajs](https://www.npmjs.com/package/kafkajs)
2. [kafka-node](https://www.npmjs.com/package/kafka-node)
3. [node-rdkafka](https://github.com/Blizzard/node-rdkafka)

For `TYS project` we were using `node-rdkafka`. `We had a few problems with missed kafka events`.

## Main Concepts and Terminology
An event records the fact that "something happened" in the world or in your business. It is also called record or message in the documentation. When you read or write data to Kafka, you do this in the form of events. Conceptually, an event has a key, value, timestamp, and optional metadata headers. Here's an example event:

- Event key: "Alice"
- Event value: "Made a payment of $200 to Bob"
- Event timestamp: "Jun. 25, 2020 at 2:06 p.m."

`Producers` are those client applications that publish (write) events to Kafka, and `consumers` are those that subscribe to (read and process) these events. In Kafka, producers and consumers are fully decoupled and agnostic of each other, which is a key design element to achieve the high scalability that Kafka is known for. For example, producers never need to wait for consumers. Kafka provides various guarantees such as the ability to process events exactly-once.

Events are organized and durably stored in `topics`. Very simplified, a topic is similar to a `folder` in a filesystem, and the events are the files in that folder. An example topic name could be "payments". Topics in Kafka are always multi-producer and multi-subscriber: a topic can have zero, one, or many producers that write events to it, as well as zero, one, or many consumers that subscribe to these events. Events in a topic can be read as often as needed—unlike traditional messaging systems, events are not deleted after consumption. Instead, you define for how long Kafka should retain your events through a per-topic configuration setting, after which old events will be discarded. Kafka's performance is effectively constant with respect to data size, so storing data for a long time is perfectly fine.

Topics are `partitioned`, meaning a topic is spread over a number of "buckets" located on different Kafka brokers. This distributed placement of your data is very important for scalability because it allows client applications to both read and write the data from/to many brokers at the same time. When a new event is published to a topic, it is actually appended to one of the topic's partitions. Events with the same event key (e.g., a customer or vehicle ID) are written to the same partition, and Kafka guarantees that any consumer of a given topic-partition will always read that partition's events in exactly the same order as they were written.

## [How does Kafka work in a nutshell?](https://kafka.apache.org/documentation/#intro_nutshell)
### Servers: 
Kafka is run as a cluster of one or more servers that can span multiple data-centers or cloud regions. Some of these servers form the storage layer, called the `brokers`.

The difference between a Kafka **broker** and a **controller** lies in their roles within the Kafka cluster, and how they manage and maintain the data and health of the Kafka system.

### 1. **Kafka Broker**

A **Kafka broker** is a server in a Kafka cluster responsible for handling the following tasks:

- **Message Storage and Delivery**: Brokers store data (messages) in topics and partitions, and serve these messages to consumers when they request them. They receive messages from producers and ensure data durability by writing them to disk.
- **Partition Leader Management**: Each broker is responsible for managing one or more partitions of a topic. In a Kafka cluster, each partition has a "leader," and the leader broker handles all read/write requests for that partition.
- **Scalability**: Kafka can scale horizontally by adding more brokers to distribute the load, both for message storage and processing.
- **Fault Tolerance**: Multiple brokers provide fault tolerance, as the replication feature ensures that data is copied across brokers, so that if one broker fails, others can still provide the data.

Each Kafka broker is identified by its unique `broker.id` and can handle many partitions for different topics.

### 2. **Kafka Controller**

The **Kafka controller** is a specialized role within the Kafka cluster. It is a **single broker** that is elected to manage cluster-level operations and maintain the health of the entire Kafka system. The controller is responsible for:

- **Leader Election**: The controller manages which brokers serve as leaders for partitions. When a leader broker fails, the controller detects the failure and triggers the election of a new leader for the affected partitions from the set of replicas.
- **Metadata Management**: The controller is responsible for managing and distributing cluster metadata (e.g., partition leader assignments) to other brokers and clients.
- **Partition Rebalancing**: When brokers are added or removed, the controller ensures partitions are reassigned or rebalanced among the available brokers.

### Key Differences:

| **Aspect**          | **Kafka Broker**                                      | **Kafka Controller**                                     |
|---------------------|-------------------------------------------------------|----------------------------------------------------------|
| **Main Role**        | Stores and serves data, handles client requests (producers/consumers). | Manages cluster-level operations, including leader election and rebalancing. |
| **Number**          | There are multiple brokers in a Kafka cluster.         | There is only one controller in the Kafka cluster at any time. |
| **Leader Management**| Each broker can be a leader for some partitions.       | The controller handles the assignment of partition leaders. |
| **Failure Handling** | Brokers store data redundantly across replicas to handle failure. | The controller detects broker failures and initiates leader elections. |
| **Election**        | Brokers do not need to be elected (each broker is independent). | The controller is elected from the brokers using ZooKeeper. |

In summary:
- A **Kafka broker** is responsible for storing, handling, and delivering messages.
- A **Kafka controller** is a single broker with the added responsibility of managing the cluster’s overall health, including leader election and partition rebalancing.

If the controller broker fails, another broker is automatically elected to take over the controller role, ensuring the cluster remains operational.

### Clients:
They allow you to write distributed applications and microservices that read, write, and process streams of events in parallel, at scale, and in a fault-tolerant manner even in the case of network problems or machine failures.

Producers are those client applications that publish (write) events to Kafka, and consumers are those that subscribe to (read and process) these events.
