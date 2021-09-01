---
date: 2000-01-01T00:00:00+00:00
title: Services
author: bradrydzewski
weight: 5
toc: true
description: |
  Configure service containers.
---

Drone supports launching detached service containers as part of your pipeline. The typical use case for services is when your unit tests require a running redis server, for example:

{{< highlight text "linenos=table,hl_lines=5-7" >}}
kind: pipeline
type: kubernetes
name: default

services:
- name: cache
  image: redis
{{< / highlight >}}

Service containers share the same network as your pipeline steps and can be access at the localhost address. In our previous example, the redis container can be accessed from the pipeline at `tcp://127.0.0.1:6379`

{{< highlight text "linenos=table,hl_lines=9" >}}
kind: pipeline
type: kubernetes
name: default

steps:
- name: ping
  image: redis
  commands:
  - redis-cli -h 127.0.0.1 ping

services:
- name: cache
  image: redis
{{< / highlight >}}

It is important to note the service container exit code is ignored, and a non-zero exit code does not fail the overall pipeline. Drone expects service containers to exit with a non-zero exit code, since they often need to be killed after the pipeline completes.

# Detached Steps

Services can also be defined directly in the pipeline, as detached pipeline steps. This can be useful when you need direct control over when the service is started, relative to other steps in your pipeline.

{{< highlight text "linenos=table,hl_lines=8" >}}
kind: pipeline
type: kubernetes
name: default

steps:
- name: cache
  image: redis
  detach: true

- name: ping
  image: redis
  commands:
  - redis-cli ping
{{< / highlight >}}

# Common Problems

This section highlights some common problems that users encounter when configuring services. If you continue to experience issues please also check the faq. You might also want to compare your yaml to our example service configurations.

## Initialization

It is important to remember that after a container is started, the software running inside the container (e.g. redis) takes time to initialize and begin accepting connections. There are two ways to handle this add a health check (prefered) or add a sleep.

### Health check

Using a commandline tool to check if a service is up and running. This is a common idiom for webservers and also docker images. Below is an example for Mysql that uses the `mysqladmin` tool to check if the mysql server is running, then runs a sql command. 

{{< highlight patch >}}
kind: pipeline
type: kubernetes
name: default

steps:
  - name: mysql healthcheck
    image: mysql:5.7
    commands:
      - while ! mysqladmin ping -h mysql-server -u drone -pdrone --silent; do sleep 1; done
      - mysql -h mysql-server -u drone -pdrone -e "CREATE TABLE IF NOT EXISTS drone_db.pipelines (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50) NOT NULL);"

services:
  - name: mysql-server
    image: mysql:5.7
    environment:
      MYSQL_ALLOW_EMPTY_PASSWORD: yes
      MYSQL_DATABASE: drone_db
      MYSQL_USER: drone
      MYSQL_PASSWORD: drone
{{< / highlight >}}

Here are some example health checks using http requests [here](https://healthchecks.io/docs/bash/). If however you are unable to craft a health check you can implement a sleep (see below).

### Sleep

Be sure to give the service adequate time to initialize before attempting to connect. A naive solution is to use the sleep command.

{{< highlight patch >}}
kind: pipeline
type: kubernetes
name: default

steps:
  - name: ping
    image: redis
    commands:
+   - sleep 5
    - redis-cli -h cache ping

services:
  - name: cache
    image: redis
{{< / highlight >}}
