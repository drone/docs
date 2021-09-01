---
date: 2000-01-01T00:00:00+00:00
title: Installation
author: bradrydzewski
weight: 2
toc: true
steps: true
description: |
  Install the runner using Docker.
---

This article explains how to install the digitalocean runner on Linux using Docker. The digitalocean runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-digitalocean).

# Download

Install Docker and pull the public image:

```
$ docker pull drone/drone-runner-digitalocean
```

# Upload Keys

Create an ssh keypair and add to Digital Ocean. The runner will use the ssh key for remote authentication and access to your droplets.

# Configuration

The runner is configured using environment variables. This article references the below configuration options. See [Configuration]({{< relref "configuration/reference" >}}) for a complete list of configuration options.

* __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.

* __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.

* __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

* __DRONE_PUBLIC_KEY_FILE__
  : provides the public key used for remote ssh access to the machine. This public key must also be added to your digital ocean account.

* __DRONE_PRIVATE_KEY_FILE__
  : provides the private key used for remote ssh access to the machine.

# Installation

The below command creates a container and starts the runner. _Remember to replace the environment variables below with your Drone server details._

{{< highlight handlebars "linenos=table" >}}
docker run -d \
  --volume=/path/on/host/id_rsa:/path/in/container/id_rsa \
  --volume=/path/on/host/id_rsa.pub:/path/in/container/id_rsa.pub \
  --env=DRONE_RPC_PROTO=https \
  --env=DRONE_RPC_HOST={{DRONE_RPC_HOST}} \
  --env=DRONE_RPC_SECRET={{DRONE_RPC_SECRET}} \
  --env=DRONE_PUBLIC_KEY_FILE={{/path/in/container/id_rsa.pub}} \
  --env=DRONE_PRIVATE_KEY_FILE={{/path/in/container/id_rsa}} \
  --publish=3000:3000 \
  --restart always \
  --name runner \
  drone/drone-runner-digitalocean
{{< / highlight >}}

# Verification

Use the `docker logs` command to view the logs and verify the runner successfully established a connection with the Drone server.

```
$ docker logs runner

INFO[0000] starting the server
INFO[0000] successfully pinged the remote server 
```
