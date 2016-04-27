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

In order to configure your build you must include a `.drone.yml` file in the root of your repository. The Yaml is a superset of the Docker compose file format. Below is an example Yaml configuration file with a single build step:

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
```

You can break your build into multiple named steps (see below example). Each step executes in a separate Docker container with shared disk access to your project workspace. Note that you can mix custom command and plugin steps.

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

  notify:
    image: slack
    channel: developers
    username: drone
```

Note that the above step names are completely arbitrary. You can call them whatever you like.

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

Drone automatically clones your repository into a local volume that is mounted into each Docker container. This volume is generally referred to as the workspace. The workspace is available to all steps in your build process, including plugins and service containers.

```
git clone --depth=50 --recusive=true \
    https://github.com/octocat/hello-world.git \
    /drone/src/github.com/octocat/hello-world

git checkout 7fd1a60
```

# Commands

Drone previously cloned your source code into the workspace. Drone mounts the workspace into your build containers (golang) and executes bash commands inside your build container, using the root of your repository as the working directory.

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
```

There is no magic here. Drone converts the above Yaml into a simple shell script that gets executed as the entrypoint to your build container. The above Yaml file is roughly translated into the below shell script:

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

# Plugins

Drone supports a publish, deployment and notification capabilities through external plugins. Plugins are Docker containers that are automatically downloaded, attach to your build, and perform a specific task.

Example Yaml configuration triggers a Heroku deployment:

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test

  deploy:
    image: heroku
    app: octokit
```

Example Yaml configuration triggers a Slack notification:

```yaml
script:
  ...

  notify:
    image: slack
    channel: developers
    username: drone
```

# Filters

Drone supports limiting execution of steps at runtime using filters. You can define filters using the `when` clause. This can be useful when you want to limit deployment steps to specific branches or notification steps to specific build statuses.

Example Yaml triggers a Heroku deployment based on branch:

```yaml
script:
  ...

  prod:
    image: heroku
    app: octokit
    when:
      branch: master

  stage:
    image: heroku
    app: octokit
    when:
      branch: feature/*
```

Example Yaml triggers a Slack notification for failures only:

```yaml
script:
  ...

  slack:
    channel: dev
    username: drone
    when:
      status: [ failure ]
```
