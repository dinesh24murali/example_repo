# How to use Traefik with NodeJS

This example uses traefik in a docker container.

### Step 1:

Pull the docker image:
```bash
docker pull traefik
```

### Step 2:

Run traefik docker image:

```bash
docker run -d -p 8080:8080 -p 80:80 \
    -v $PWD/traefik/traefik.yml:/etc/traefik/traefik.yml traefik:v3.1
```