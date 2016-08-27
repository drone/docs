+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Quick Start"
weight = 1
toc = true
break = true

[menu.main]
	parent="usage"

+++

# Overview

Configure your build by placing a `.drone.yml` file in the root of your repository. The `.drone.yml` is a superset of the [Docker Compose](https://docs.docker.com/compose/) file format. Below is an example configuration file with a single build step that runs tests against a Postgres service container.

```yaml
pipeline:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test

services:
  postgres:
    image: postgres:9.4.5
    environment:
      - POSTGRES_USER=myapp
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

For private repositories credentials can be configured by secrets
`REGISTRY_USERNAME`, `REGISTRY_PASSWORD` and `REGISTRY_EMAIL`, example:

```
# following adds contents of password.txt to secret REGISTRY_PASSWORD
# and REGISTRY_PASSWORD will be when pulling gcr.io/me/repo:tag
$ drone secret add --image=gcr.io/me/repo:tag my/repo REGISTRY_PASSWORD @password.txt
```

# Cloning

Drone automatically clones your repository into a local volume that is mounted into each Docker container. This volume is generally referred to as the workspace. The workspace is available to all steps in your build process, including plugins and service containers.

```
git clone --depth=50 --recusive=true \
    https://github.com/octocat/hello-world.git \
    /drone/src/github.com/octocat/hello-world

git checkout 7fd1a60
```

See [Cloning](../cloning) for more details.

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

See [Services](../services) for more details.

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
    app: foo.com
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

# Constraints

Drone gives you the ability to conditionally limit the execution of build steps at runtime. The below example limits execution of Heroku plugin steps based on branch:

```yaml
pipeline:
  ...

  prod:
    image: heroku
    app: foo.com
    when:
      branch: master

  stage:
    image: heroku
    app: dev.foo.com
    when:
      branch: feature/*
```

See [Constraints](../constraints) for a more detailed overview of how this works.

# Failures

Drone uses the container exit code to determine the success or failure status of a build. Non-zero exit codes fail the build and cause the pipeline to immediately exit.

There are use cases for executing pipeline steps on failure, such as sending notifications for failed builds. Use the status constraint to override the default behavior and execute steps even when the build status is failure:

```
pipeline:
  build:
    ...

  notify:
    image: slack
    when:
      status: [ success, failure ]
```
