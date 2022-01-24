---
date: 2000-01-01T00:00:00+00:00
title: HA
author: tphoney
weight: 21
toc: true
description: |
  Configure Drone server HA.
---

<div class="alert">
HA (high availability) for drone is in Beta and may not be suitable for production workloads. Furthermore this is a community effort and is not subject to support services or service level agreements at this time.
</div>

Drone is a now capable of being a highly available service that is resilient to failures.

# Overview

![HA diagram](/ha_server.png)

There are two new components that are required to make drone highly available.

- The [HAProxy](https://www.haproxy.org/) load balancer. Used to load balance the drone server UI, and also to balance requests from the drone runners to the drone servers. The sample configuration in this article uses HAProxy, however, any standard Load Balancer or reverse proxy should work.
- The [Redis](https://redis.io/) for queues and pub/sub. Used for runner events, log streaming, build cancelation events and finally the build queue itself.

# Configuration

Here we will take a look at a basic configuration for HA and how to configure the load balancer and redis.

## Drone server configuration

- makes the assumption we are running postgres db on the same host as the drone servers.
- makes the assumption we are running redis db on the same host as the drone servers. This is the trigger for Drone to run in HA mode.

For drone server 1

```bash
DRONE_SERVER_HOST=localhost:8081
DRONE_SERVER_PORT=:8081
DRONE_SERVER_PROTO=http
DRONE_DATABASE_DRIVER=postgres
DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@localhost:5432/drone?sslmode=disable
DRONE_REDIS_CONNECTION=redis://localhost:6379
```

For drone server 2

```bash
DRONE_SERVER_HOST=localhost:8082
DRONE_SERVER_PORT=:8082
DRONE_SERVER_PROTO=http
DRONE_DATABASE_DRIVER=postgres
DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@localhost:5432/drone?sslmode=disable
DRONE_REDIS_CONNECTION=redis://localhost:6379
```

## HAProxy

HA Proxy can work on layer 4 or layer 7 (OSI model) mode. Both modes can be used. Example `haproxy.cfg` files follow. HA proxy can be start with the following `haproxy -f haproxy.cfg` command.

Now, the Drone frontend app can be opened at http://localhost:8080. 

### Layer 4 mode

```bash
frontend drone
    bind *:8080
    default_backend drone_servers
    timeout client 120s

backend drone_servers
    balance roundrobin
    timeout connect 5s
    timeout server 120s
    server server1 localhost:8081
    server server2 localhost:8082
```

### Layer 7 mode

Note that in layer 7 mode HTTP responses will have an additional header `X-App-Server` that tells which server processed the request.

```bash
defaults
    mode http
    timeout connect 5s
    timeout client 120s
    timeout server 120s
    maxconn 128

frontend drone
    bind *:8080
    default_backend drone_servers

backend drone_servers
    balance roundrobin
    server server1 localhost:8081
    server server2 localhost:8082
    http-response add-header X-App-Server %b/%s

```

# Docker compose example

***Note: This is for demonstration purposes only, and not for production.***

Running all the containers in a single machine does not make sense, this example is provided as an example of a HA deployment.

You will need to create these two files in a directory and subsitiute in correct values to the docker-compose.yml file.

Start the stack via `docker-compose up`.

## `haproxy.cfg` file

```bash
defaults
    mode http
    timeout connect 5s
    timeout client 120s
    timeout server 120s
    maxconn 128

frontend drone
    bind *:8080
    default_backend drone_servers

backend drone_servers
    balance roundrobin
    server server1 dronea:8081
    server server2 droneb:8082
    http-response add-header X-App-Server %b/%s
```

## `docker-compose.yml` file

```yaml
version: "3.8"
services:
    dronea:
        image: drone/drone:latest
        depends_on: 
            db:
                condition: service_healthy
        environment:
        - DRONE_SERVER_HOST=localhost:8081
        - DRONE_SERVER_PORT=:8081
        - DRONE_SERVER_PROTO=http
        - DRONE_SERVER_PROXY_HOST=${DRONE_SERVER_PROXY_HOST}
        - DRONE_SERVER_PROXY_PROTO=https
        - DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6
        - DRONE_COOKIE_SECRET=e8206356c843d81e05ab6735e7ebf075
        - DRONE_COOKIE_TIMEOUT=720h
        - DRONE_GITHUB_CLIENT_ID=${DRONE_GITHUB_CLIENT_SECRET}
        - DRONE_GITHUB_CLIENT_SECRET=${DRONE_GITHUB_CLIENT_SECRET}
        - DRONE_LOGS_DEBUG=true
        - DRONE_CRON_DISABLED=true
        - DRONE_DATABASE_DRIVER=postgres
        - DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@db:5432/drone?sslmode=disable
        - DRONE_REDIS_CONNECTION=redis://redis-server:6379
    droneb:
        image: drone/drone:latest
        environment:
        - DRONE_SERVER_HOST=localhost:8082
        - DRONE_SERVER_PORT=:8082
        - DRONE_SERVER_PROTO=http
        - DRONE_SERVER_PROXY_HOST=${DRONE_SERVER_PROXY_HOST}
        - DRONE_SERVER_PROXY_PROTO=https
        - DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6
        - DRONE_COOKIE_SECRET=e8206356c843d81e05ab6735e7ebf075
        - DRONE_COOKIE_TIMEOUT=720h
        - DRONE_GITHUB_CLIENT_ID=${DRONE_GITHUB_CLIENT_SECRET}
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
        - DRONE_RPC_HOST=proxy:8080
        - DRONE_SERVER_PORT=:8080
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
            - "8080:8080"
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
