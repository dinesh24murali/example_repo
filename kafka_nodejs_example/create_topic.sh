#/bin/bash

docker exec -i broker bash /opt/kafka/bin/kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092
echo "topic test-topic was create"
