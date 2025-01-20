# Single node elastic search

0. Optional step
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
It will look something like this:
```
password
```

5. Get enrolment token
```bash
docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-create-enrollment-token -s kibana
```
It will look something like this:
```
eyJ2ZXIiOiI4LjE0LjA0ZDAiLCJrZXkiOiI4Q21rY1pRQnR1MXpxVTJ1VWhOdTpBR3E3QnIyc1RsT01sdE5hQ0YtTEFRIn0=
```

6. Create the following directories from where you run the command.
```
./logstash/credentials/ca/
./logstash/pipeline
./logstash/settings
```

7. Get ssl certificate

```bash
docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt ./logstash/credentials/ca/ca.crt
```

8. Run kibana

```bash
docker run --name kib01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.17.0
```

The kibana container will print out a URL that you can access through the browser. It will look like this:

```
i Kibana has not been configured.

Go to http://0.0.0.0:5601/?code=234301 to get started.
```

9. To run logstash you need to make some changes to a few configurations. You need to get the IP address of the elastic search container. Use the following command:

```bash
docker inspect es01 | grep "IPAddress"
```
It will printout something like the following:
```
"SecondaryIPAddresses": null,
"IPAddress": "",
"IPAddress": "111.11.1.1",
```

10. Open `./logstash/pipeline/logstash.conf` file, and update the output section where we have the password. Also update the IP address of the elastic container under the hosts section:
```
output {
  elasticsearch {
    index => "logstash-%{+YYYY.MM.dd}"
    hosts => ["https://111.11.1.1:9200"] <------------------ Change the IP address here
    user => "elastic"
    password => "password" <------------------ Change the password
    ssl_enabled => true
    cacert => "/usr/share/logstash/certs/ca/ca.crt"
  }
}
```

11. This is the command the start logstash. Remember to run this command from the folder where we previously created the logstash folder. Update the IP address of elastic in the `ELASTIC_HOSTS` environment variable, and update the elasticsearch password in the `ELASTIC_PASSWORD` environment variable.
```bash
docker run --rm -it --net elastic \
-v ./logstash/pipeline/:/usr/share/logstash/pipeline/ \
-v ./logstash/settings/logstash.yml:/usr/share/logstash/config/logstash.yml \
-v ./logstash/credentials/:/usr/share/logstash/certs \
-p 5044:5044/udp \
docker.elastic.co/logstash/logstash:8.17.0
```

# Adding data to elastic search

Add the following environment variable to the terminal
```bash
export ELASTIC_PASSWORD="password"
```

Remember to run the following commands from the same directory where the ssl certificate is present. If you were following the above steps the ssl certificate will be present in `./logstash/credentials/ca`

1. Create a new index called `books`
```bash
curl -X PUT --cacert ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books?pretty
```
2. Add data to the `book` index:
```bash
curl -X POST --cacert ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_doc?pretty -H 'Content-Type: application/json' -d'
{
  "name": "Snow Crash",
  "author": "Neal Stephenson",
  "release_date": "1992-06-01",
  "page_count": 470
}
'
```
4. Bulk add data to the `book` index:
```bash
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
```
5. Get the field mappings
```bash
curl -X GET --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_mapping?pretty
```
6. Get the field mappings
```bash
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
```
7. Query elastic search:
```bash
curl -X GET --cacert ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_search?pretty -H 'Content-Type: application/json' -d'
{
  "query": {
    "match": {
      "name": "brave"
    }
  }
}
'
```

8. Update a document:
Remember to replace `<ID of document>` with the ID of the document.
```bash
curl -X POST --cacert ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200/books/_update/<ID of document> -H 'Content-Type: application/json' -d'
{
  "doc": {
    "field_to_update": "new_value"
  }
}
'
```