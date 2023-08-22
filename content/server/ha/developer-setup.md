---
date: 2000-01-01T00:00:00+00:00
title: Local Development
author: tphoney
weight: 21
toc: true
description: |
  HA setup for developers.
---

This is purely for development purposes only. All the services of the HA setup are run on a single machine and is only suitable for local testing.

You will need to create a directory that contains a `docker-compose.yml` file substituting in correct values, and a `haproxy.cfg` file.

## `docker-compose.yml` file

- DRONE_SERVER_PROXY_HOST is used to set the proxy host for the drone server. eg using ngrok. Explained in greater detail below.
- DRONE_GITHUB_CLIENT_ID and DRONE_GITHUB_CLIENT_SECRET are used to set the github client id and secret.

```yaml {linenos=table}
version: "3.8"
services:
    dronea:
        image: drone/drone:latest
        depends_on: 
            db:
                condition: service_healthy
        environment:
        - DRONE_SERVER_HOST=localhost
        - DRONE_SERVER_PROTO=http
        - DRONE_SERVER_PROXY_HOST=${DRONE_SERVER_PROXY_HOST}
        - DRONE_SERVER_PROXY_PROTO=https
        - DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6
        - DRONE_COOKIE_SECRET=e8206356c843d81e05ab6735e7ebf075
        - DRONE_COOKIE_TIMEOUT=720h
        - DRONE_GITHUB_CLIENT_ID=${DRONE_GITHUB_CLIENT_ID}
        - DRONE_GITHUB_CLIENT_SECRET=${DRONE_GITHUB_CLIENT_SECRET}
        - DRONE_LOGS_DEBUG=true
        - DRONE_CRON_DISABLED=true
        - DRONE_DATABASE_DRIVER=postgres
        - DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@db:5432/drone?sslmode=disable
        - DRONE_REDIS_CONNECTION=redis://redis-server:6379
    droneb:
        image: drone/drone:latest
        environment:
        - DRONE_SERVER_HOST=localhost
        - DRONE_SERVER_PROTO=http
        - DRONE_SERVER_PROXY_HOST=${DRONE_SERVER_PROXY_HOST}
        - DRONE_SERVER_PROXY_PROTO=https
        - DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6
        - DRONE_COOKIE_SECRET=e8206356c843d81e05ab6735e7ebf075
        - DRONE_COOKIE_TIMEOUT=720h
        - DRONE_GITHUB_CLIENT_ID=${DRONE_GITHUB_CLIENT_ID}
        - DRONE_GITHUB_CLIENT_SECRET=${DRONE_GITHUB_CLIENT_SECRET}
        - DRONE_LOGS_DEBUG=true
        - DRONE_CRON_DISABLED=true
        - DRONE_DATABASE_DRIVER=postgres
        - DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@db:5432/drone?sslmode=disable
        - DRONE_REDIS_CONNECTION=redis://redis-server:6379
        depends_on: 
        - dronea
    runner:
        image: drone/drone-runner-docker:latest
        environment:
        - DRONE_RPC_HOST=proxy
        - DRONE_RPC_PROTO=http
        - DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6
        - DRONE_TMATE_ENABLED=true
        volumes:
        - /var/run/docker.sock:/var/run/docker.sock
        depends_on: 
        - proxy
        - droneb 
    redis-server:
      image: redis
    proxy:
        image: haproxy
        ports:
            - "8080:80"
        volumes:
        - ./haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg
        depends_on: 
        - dronea
    db:
        image: postgres
        environment:
        - POSTGRES_USER=postgres
        - POSTGRES_PASSWORD=postgres
        - POSTGRES_DB=drone
        healthcheck:
            test: ["CMD-SHELL", "pg_isready -U postgres"]
            interval: 5s
            timeout: 5s
            retries: 5
```

## `haproxy.cfg` file

```bash {linenos=table}
defaults
    mode http
    timeout connect 5s
    timeout client 120s
    timeout server 120s
    maxconn 128

frontend drone
    bind *:80
    default_backend drone_servers

backend drone_servers
    balance roundrobin
    server server1 dronea:80
    server server2 droneb:80
    http-response add-header X-App-Server %b/%s
```

## Using Ngrok as a proxy for the drone server

Ngrok allows us to send the webhooks from Github to our local Drone setup.

Follow the guide here `https://dashboard.ngrok.com/get-started/setup`

Run Ngrok against port 8080 it will run in the foreground.

```
ngrok http 8080
```

Take note of the forwarding hostname this is your DRONE_SERVER_PROXY_HOST EG

```
Forwarding    http://c834c33asdde.ngrok.io -> http://localhost:8080
```

You can now use this ngrok hostname in your `docker-compose.yml` file for the `DRONE_SERVER_PROXY_HOST`. NB do not include http/https.

## Starting the services

Start the stack via `docker-compose up` and then open the drone UI at `http://localhost:8080`.
