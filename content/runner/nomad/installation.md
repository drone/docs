---
date: 2000-01-01T00:00:00+00:00
title: Installation
author: bradrydzewski
weight: 2
description: |
  Install the runner using Docker.
---

This article explains how to install the Nomad. The Nomad runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-nomad).
# Download

Install Docker and pull the public image:

```
$ docker pull drone/drone-runner-nomad:latest
```

# Configuration

The Nomad runner is configured using environment variables. This article references the below configuration options. See [Configuration]({{< relref "configuration/reference" >}}) for a complete list of configuration options.

* __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.

* __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.

* __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

* __DRONE_JOB_DATACENTER__
  : provides the datacenter used when scheduling Nomad jobs. The value is optional and defaults to dc1 if unspecified.

* __DRONE_JOB_REGION__
  : provides the region used when scheduling Nomad jobs.

* __DRONE_JOB_NAMESPACE__
  : provides the namespace used when scheduling Nomad jobs.

* __NOMAD_TOKEN__
  : provides the Nomad token.

* __NOMAD\_*__
  : the system supports all NOMAD_ client configuration parameters as defined in the Nomad client documentation.

# Installation

The below command creates the a container and start the Docker runner. _Remember to replace the environment variables below with your Drone server details._

```
$ docker run -d \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e DRONE_RPC_PROTO=https \
  -e DRONE_RPC_HOST=drone.company.com \
  -e DRONE_RPC_SECRET=super-duper-secret \
  -e NOMAD_ADDR=http://nomad.company.com:4646 \
  -e NOMAD_TOKEN=super-secret-token \
  --restart always \
  --name runner \
  drone/drone-runner-nomad:latest
```

# Verification

Use the `docker logs` command to view the logs and verify the runner successfully established a connection with the Drone server.

```
$ docker logs runner

INFO[0000] starting the server
INFO[0000] successfully pinged the remote server 
```