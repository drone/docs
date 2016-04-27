+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Overview"
weight = 1
menu = "usage"
toc = true
break = true
+++

# Overview

In order to configure your build you must include a `.drone.yml` file in the root of your repository. This section provides a brief overview of the Yaml configuration file and build process. This is a simple Yaml configuration file with a single step:

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
```

You can break your build into multiple named steps (see below example). Each step executes in a separate Docker container with shared disk access to your project workspace.

```yaml
script:
  backend:
    image: golang
    commands:
      - go get
      - go build
      - go test

  frontend:
    image: node:6
    commands:
      - npm install
      - npm test
```

# Activate

Before we get started you need to login to Drone and activate your repository. When you activate your repository Drone automatically adds post-commit hooks with your version control system (ie GitHub) to trigger builds when you push code or open pull requests.

# Images

Drone executes your build inside an ephemeral Docker image. This means you don't have to setup or install any repository dependencies on your host machine. Use any valid Docker image in any Docker registry as your build environment.

```yaml
script:
  build:
    image: golang:1.6
```

# Cloning

Drone automatically clones your repository into a local volume that is mounted into each Docker container. This workspace is available to all steps in your build process, including plugins and service containers.

```
git clone --depth=50 --recusive=true \
    https://github.com/octocat/hello-world.git \
    /drone/src/github.com/octocat/hello-world

git checkout 7fd1a60
```

# Scripts

Drone previously cloned your source code into the project workspace. Drone mounts the build workspace into your build containers (golang) and executes bash commands inside your build container.

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
```

There is no magic here. Drone converts the above Yaml into a simple shell script that gets executed as the entrypoint to your build container. The above Yaml file gets translated into the below shell script:

```sh
#!/bin/sh
set -e

go get
go build
go test
```

# Services

Drone supports launching service containers as part of the build process. This can be very helpful when your unit tests require database access, for example. Service containers share the same network (ie localhost) as your build containers. 

Example Yaml configuration using a Postgres database:

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test

services:
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=mysecretpassword
```

# Deployments

Drone supports a large number of publish and deployment capabilities through external plugins. Plugins are Docker containers that are automatically downloaded, attach to your build, and automatically execute publish or deployment tasks.

Example Yaml configuration triggers a Heroku deployment:

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
  heroku:
    app: octokit
    force: false
    when:
      branch: master
```

# Notifications

Drone also supports notification options through external plugins. Notification plugins are Docker containers that are automatically downloaded, attach to your build, and automatically trigger notifications.

Example Yaml configuration triggers a Slack notification:

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
  slack:
    channel: dev
    username: drone
    when:
      status: [ success, failure ]
```

# Getting Help

For help troubleshooting failed builds please use [Stackoverflow](https://stackoverflow.com). The Stackoverflow community will be able to answer questions unique to your programming language and technology stack that the Drone maintainers are unqualified to answer.

For all other questions or issues please use the community [chat room](https://gitter.im/drone/drone) for support.
