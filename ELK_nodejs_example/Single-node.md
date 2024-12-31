# Single node elastic search

```bash
sudo sysctl -w vm.max_map_count=262144
```

1. Create a new docker network.

```bash
docker network create elastic
```

2. Pull the Elasticsearch Docker image.

```bash
docker pull docker.elastic.co/elasticsearch/elasticsearch:8.17.0
```

3. Start an Elasticsearch container.

```bash
docker run --name es01 --net elastic -p 9200:9200 -it -m 1GB docker.elastic.co/elasticsearch/elasticsearch:8.17.0
```

4. Reset elastic user password
```bash
docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic
```

SAyYi6KnrdZLJpy*jURT

5. Get enrolment token
```bash
docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-create-enrollment-token -s kibana
```
eyJ2ZXIiOiI4LjE0LjAiLCJhZHIiOlsiMTcyLjI5LjAuMjo5MjAwIl0sImZnciI6ImNiYjE1NzljNzE2YmRhN2U4ZDcxOTMyN2U3NThhY2EwMWQ2NDFkOTZjZDA0ZDE1OGJlYzA0NjBiMDQ1OThiM2EiLCJrZXkiOiJjWGNGR3BRQmpqcFY2Y2hCbWViNDpQUm5YenhuSFRzV09CalROV0VZSmdnIn0=

6. Add environment variable
```bash
export ELASTIC_PASSWORD="SAyYi6KnrdZLJpy*jURT"
```

7. Get ssl certificate

```bash
docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt ./logstash/credentials/ca/ca.crt
```

8. Run kibana

```bash
docker run --name kib01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.17.0
```

9. Run logstash
```bash
docker run --rm -it --net elastic -v ./logstash/pipeline/:/usr/share/logstash/pipeline/ -v ./logstash/settings/logstash.yml:/usr/share/logstash/config/logstash.yml -v ./logstash/credentials/:/usr/share/logstash/certs -e ELASTIC_USER="elastic" -e ELASTIC_PASSWORD="SAyYi6KnrdZLJpy*jURT" -e ELASTIC_HOSTS="https://172.29.0.2:9200" -e xpack.monitoring.enabled=false -e NODE_NAME="logstash" -p 5044:5044/udp docker.elastic.co/logstash/logstash:8.17.0
```

curl -X PUT --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books?pretty


curl -X POST --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_doc?pretty -H 'Content-Type: application/json' -d'
{
  "name": "Snow Crash",
  "author": "Neal Stephenson",
  "release_date": "1992-06-01",
  "page_count": 470
}
'


curl -X POST --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/_bulk?pretty -H 'Content-Type: application/json' -d'
{ "index" : { "_index" : "books" } }
{"name": "Revelation Space", "author": "Alastair Reynolds", "release_date": "2000-03-15", "page_count": 585}
{ "index" : { "_index" : "books" } }
{"name": "1984", "author": "George Orwell", "release_date": "1985-06-01", "page_count": 328}
{ "index" : { "_index" : "books" } }
{"name": "Fahrenheit 451", "author": "Ray Bradbury", "release_date": "1953-10-15", "page_count": 227}
{ "index" : { "_index" : "books" } }
{"name": "Brave New World", "author": "Aldous Huxley", "release_date": "1932-06-01", "page_count": 268}
{ "index" : { "_index" : "books" } }
{"name": "The Handmaids Tale", "author": "Margaret Atwood", "release_date": "1985-06-01", "page_count": 311}
'


curl -X GET --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_mapping?pretty


curl -X PUT --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/my-explicit-mappings-books?pretty -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "dynamic": false,  
    "properties": {  
      "name": { "type": "text" },
      "author": { "type": "text" },
      "release_date": { "type": "date", "format": "yyyy-MM-dd" },
      "page_count": { "type": "integer" }
    }
  }
}
'

curl -X GET --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_search?pretty

curl -X GET --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_search?pretty -H 'Content-Type: application/json' -d'
{
  "query": {
    "match": {
      "name": "brave"
    }
  }
}
'
