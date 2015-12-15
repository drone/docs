+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Quick Start"
weight = 1
menu = "usage"
toc = true
+++

# Overview

In order to configure your build you must include a `.drone.yml` file in the root of your repository. This section provides a brief overview of the configuration file and build process.

Example .drone.yml configuration:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test
```

# Hooks

Once activated, every commit and pull request automatically send a hook from your version control system (ie GitHub) to Drone. These hooks instruct Drone to execute a new build.

When Drone receives a hook it fetches the `.drone.yml` from your repository and uses this as a blueprint for build execution. For the purposes of this tutorial let's assume your repository uses the following configuration:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test
```

# Cloning

Hooks specify commit details including branch and commit hash. Drone will automatically clone and checkout the commit into the build workspace:

```
git clone --depth=50 --recusive=true \
    https://github.com/octocat/hello-world.git \
    /drone/src/github.com/octocat/hello-world

git checkout 7fd1a60
```

# Images

Drone executes your build inside an ephemeral Docker image. This means you don't have to setup or install any repository dependencies on your host machine. Use any valid Docker image in any Docker registry as your build environment.

Example .drone.yml configuration uses the official Go Docker image:

```yaml
---
build:
  image: golang
```

# Commands

Drone previously cloned your source code into the build workspace. The build workspace is mounted into your Docker container at runtime as a volume. This means your code is cloned **outside** of the build container but your build commands and run **inside** of the build container.

Drone executes the following bash commands inside your build container:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test
```

# Services

Drone supports launching separate, ephemeral Docker containers as part of the build process. This is useful, for example, if you require a database for running your unit tests.

Example .drone.yml configuration with a Postgres database:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test

compose:
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=mysecretpassword
```

# Deployments

Drone supports a large number of publish and deployment capabilities through external plugins. Plugins are Docker containers that are automatically downloaded, attach to your build, and execute a very specific publish or deployment task.

Example .drone.yml configuration with the Docker publish plugin:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test

publish:
  docker:
    username: octocat
    password: password
    email: octocat@github.com
    repo: octocat/hello-world
```

First Drone runs your build commands inside the Golang container:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test
```

Drone executes publish and deployment plugins upon successful completion of the build step. Plugins are executed in separate Docker containers but have access to your build workspace. This means any files created and stored in the `/drone` workspace are available to plugins.

The Docker plugin in our example runs `docker build` and `docker publish` after the build step successfully completes using the configuration parameters in the .drone.yml file:

```yaml
---
publish:
  docker:
    username: octocat
    password: password
    email: octocat@github.com
    repo: octocat/hello-world
```

# Local Testing

Download the [command line tools](/devs/cli) to build and test your code locally inside a Docker environment using the exact same build process as Drone. You should think of your `.drone.yml` file as a `docker-compose.yml` alternative that is optimized for repeatable, local testing.

Command to execute a local build from the command line:

```
drone exec
```

# Getting Help

For help troubleshooting **failed builds** please use [Stackoverflow](https://stackoverflow.com). The Stackoverflow community will be able to answer questions unique to your programming language and technology stack that the Drone maintainers are unqualified to answer.

For all other questions or issues please use the community [chat room](https://gitter.im/drone/drone) for support.
