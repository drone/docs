---
date: 2000-01-01T00:00:00+00:00
title: Install On Linux
title_in_sidebar: On Linux
author: bradrydzewski
weight: 1

toc: true
steps: true
aliases:
- /runners/docker/install/linux-amd64
- /runners/docker/install/linux-arm
- /runners/docker/install/linux-arm64
description: |
  Install the runner using Docker.
---

This article explains how to install the Docker runner on Linux. The Docker runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-docker), and is available for the following architectures:

* amd64
* arm
* arm64

# Download

Install Docker and pull the public image:

```
$ docker pull drone/drone-runner-docker:1
```

# Configuration

The Docker runner is configured using environment variables. This article references the below configuration options. See [Configuration]({{< relref "../configuration/reference" >}}) for a complete list of configuration options.

* __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.

* __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.

* __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

# Installation

The below command creates a container and starts the Docker runner. _Remember to replace the environment variables below with your Drone server details._

```
{{< highlight handlebars "linenos=table" >}}
docker run --detach \
  --volume=/var/run/docker.sock:/var/run/docker.sock \
  --env=DRONE_RPC_PROTO={{DRONE_RPC_PROTO}} \
  --env=DRONE_RPC_HOST={{DRONE_RPC_HOST}} \
  --env=DRONE_RPC_SECRET={{DRONE_RPC_SECRET}} \
  --env=DRONE_RUNNER_CAPACITY={{DRONE_RUNNER_CAPACITY}} \
  --env=DRONE_RUNNER_NAME={{DRONE_RUNNER_NAME}} \
  --publish=3000:3000 \
  --restart=always \
  --name=runner \
  drone/drone-runner-docker:1
{{< / highlight >}}
```

# Verification

Use the `docker logs` command to view the logs and verify the runner successfully established a connection with the Drone server.

```
$ docker logs runner

INFO[0000] starting the server
INFO[0000] successfully pinged the remote server 
```