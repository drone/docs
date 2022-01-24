---
date: 2000-01-01T00:00:00+00:00
title: Overview
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

- The load balancer [HAProxy](http://www.haproxy.org). Used to load balance the drone server UI requests, and balance requests from the drone runners to the drone servers. Any standard Load Balancer or reverse proxy should work. We use HAProxy in our [developer setup](/server/ha/developer-setup).
- The [Redis](https://redis.io/) for queues and pub/sub. Used for runner events, log streaming, build cancelation events and finally the build queue itself.

# Configuration

Here we will take a look at a basic configuration for HA and how to configure the load balancer and redis.

## Drone server configuration

For our first server {DRONE_SERVER_1_HOSTNAME_OR_IP}

```bash
DRONE_SERVER_HOST={PUBLIC_HOSTNAME_OR_IP}
DRONE_SERVER_PROTO=https
DRONE_DATABASE_DRIVER=postgres
DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@{DATABASE_HOSTNAME_OR_IP}:5432/drone?sslmode=disable
DRONE_REDIS_CONNECTION=redis://{REDIS_HOSTNAME_OR_IP}:6379
```

For our second server 2 {DRONE_SERVER_2_HOSTNAME_OR_IP}

```bash
DRONE_SERVER_HOST={PUBLIC_HOSTNAME_OR_IP}
DRONE_SERVER_PROTO=https
DRONE_DATABASE_DRIVER=postgres
DRONE_DATABASE_DATASOURCE=postgres://postgres:postgres@{DATABASE_HOSTNAME_OR_IP}:5432/drone?sslmode=disable
DRONE_REDIS_CONNECTION=redis://{REDIS_HOSTNAME_OR_IP}:6379
```

Now, the Drone frontend app can be opened at `https://{PUBLIC_HOSTNAME_OR_IP}`.

# Development setup using Docker compose

***Note: This is for development purposes, and not for production.***

{{< link "/server/ha/developer-setup" >}}
