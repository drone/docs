---
date: 2017-04-15T14:39:04+02:00
title: Services
url: services

next_steps:
  - file: mysql.md
    name: Example using MySQL
  - file: postgres.md
    name: Example using Postgres
  - file: redis.md
    name: Example using Redis

menu:
  usage:
    weight: 6
    identifier: services
    parent: usage_concepts
---

Drone provides a services section in the Yaml file used for defining service containers. The below configuration composes database and cache containers.

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

services:
  database:
    image: mysql

  cache:
    image: redis
```

Services are accessed using custom hostnames. In the above example the mysql service is assigned the hostname `database` and is available at `database:3306`.

# Configuration

Service containers generally expose environment variables to customize service startup such as default usernames, passwords and ports. Please see the official image documentation to learn more.

```diff
services:
  database:
    image: mysql
+   environment:
+     - MYSQL_DATABASE=test
+     - MYSQL_ALLOW_EMPTY_PASSWORD=yes

  cache:
    image: redis
```

# Initialization

Service containers require time to initialize and begin to accept connections. If you are unable to connect to a service you may need to wait a few seconds or implement a backoff.

```diff
pipeline:
  test:
    image: golang
    commands:
+     - sleep 15
      - go get
      - go test

services:
  database:
    image: mysql
```

# Detached Services

If you want to start a service container from the pipeline that does not block the pipeline, you can use `detach: true`. This can be useful if you want to test a docker container you have built.

```yml
pipeline:
  build:
    image: docker:latest
    entrypoint: []
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    commands:
      - "docker build -t example/image:latest -f ./Dockerfile ./"

  example_image:
    image: example/image:latest
    detach: true
    
  test: # runs parallel to example_image
    image: jamrizzi/drone-postman:latest
    collection: 'https://www.getpostman.com/collections/1d9b360b4c1d263711ff'
```
