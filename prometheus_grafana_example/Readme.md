# A sample project to demonstrate Prometheus and Grafana

This docker compose includes the following:
1. [Traefik](https://traefik.io/traefik/) - used to track requests and expose them to Prometheus
2. A sample NodeJS application running in a docker container
3. [Prometheus](https://prometheus.io/) - used to scrape data from Traefik, cAdvisor, and Node exporter
4. [cadvisor](https://github.com/google/cadvisor) - used to export data about running docker containers
5. [node-exporter](https://prometheus.io/docs/guides/node-exporter/) - used to get details like RAM, CPU usage, I/O through put, disk space, etc
6. [Grafana](https://grafana.com/) - Used to visualize the data stored in Prometheus using dashboards and graphs

The Data source configurations for Grafana is inside its own folder.

The scraping config for Prometheus is inside its own folder.

## How to run:

Step 1:
Run the following command to build the image for our nodejs app
```bash
docker compose build
```

Step 2:
Run the following to start all the services
```bash
docker compose up
```
`Traefik will run in port 80`

- Traefik will send all request to the not application

Step 3:
Open this link to trigger a few http requests so that Traefik will create some data for Prometheus.

for 500 error:
http://node.localhost/error_or_not/error

for 200:

http://node.localhost/error_or_not/sometext

http://node.localhost/

A few links:

- Access Prometheus expression browser using this [link](http://localhost:9090/graph?g0.expr=prometheus_http_requests_total%7Bcode%3D~%222.*%22%7D&g0.tab=0&g0.display_mode=lines&g0.show_exemplars=0&g0.range_input=1h)

- Access cAdvisor dashboard using this [link](http://localhost:8080/)

- Access Grafana using this [link](http://localhost:3000/)

Step 4:

- Login to Grafana using the following credentials:
```
username: admin
password: grafana
```

- Prometheus will already be added as the default data source.

Step 5:

- Click on `import dashboard`

- Open the following links to get the Dashboard id for node-exporter

https://grafana.com/grafana/dashboards/1860-node-exporter-full/

- paste that id to the text box visible in the UI. The dashboard name will auto populate

- Select Prometheus as datasource

- click on save

- The dashboard should be visible now.

- use the following links to setup the other dashboards in a similar fashion

#### Grafana dashboards:

1. Node exporter: https://grafana.com/grafana/dashboards/1860-node-exporter-full/
2. cAdvisor: https://grafana.com/grafana/dashboards/14282-cadvisor-exporter/
3. Traefik: https://grafana.com/grafana/dashboards/17346-traefik-official-standalone-dashboard/