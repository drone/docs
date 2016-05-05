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

Configure your build by placing a `.drone.yml` file in the root of your repository. The `.drone.yml` is a superset of the Docker compose file format. Below is an example configuration file with a single build step.

```yaml
pipeline:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
```

You can break your build into multiple named steps (see below example). Each step executes in a separate Docker container with shared disk access to your project workspace. Note that you can mix command execution and plugin steps.

```yaml
pipeline:
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
pipeline:
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
pipeline:
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
pipeline:
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

Drone supports publish, deployment and notification capabilities through external plugins. Plugins are Docker containers that are automatically downloaded, attach to your build, and perform a specific task.

Example Yaml configuration triggers a Heroku deployment:

```yaml
pipeline:
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
pipeline:
  ...

  notify:
    image: slack
    channel: developers
    username: drone
```

# Conditions

Drone gives you the ability to conditionally limit the execution of build steps at runtime. The below example limits execution of heroku plugin steps based on branch:

```yaml
pipeline:
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

Execute a step if the branch is `master` or `develop`:

```yaml
when:
  branch: [master, develop]
```

Execute a step if the branch is starts with `prefix/*`:

```yaml
when:
  branch: prefix/*
```

Execute a step if the build event is a `tag`:

```yaml
when:
  event: tag
```

Execute a step for all non-pull request events (default):

```yaml
when:
  event: [push, tag, deployment]
```

Execute a step for all events:

```yaml
when:
  event: [push, pull_request, tag, deployment]
```

Execute a step when the build status changes:

```yaml
when:
  status: changed
```

Execute a step when the build is passing or failing:

```yaml
when:
  status:  [ failure, success ]
```

Execute a step for a specific matrix combination:

```yaml
when:
  matrix:
    GO_VERSION: 1.4
    REDIS_VERSION: 3.0
```

Execute a step for a specific platform:

```yaml
when:
  platform: [ linux/amd64, linux/arm ]
```

Execute a step for a specific platform using wildcards:

```yaml
when:
  platform:  linux/*
```

Execute a step for deployment events matching the target environment:

```yaml
when:
  environment: production
  event: deployment
```
