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

This article explains how to install the MacStadium runner on Linux using Docker. The MacStadium runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-macstadium).

# Download

Install Docker and pull the public image:

```
$ docker pull drone/drone-runner-macstadium
```

# Networking

The MacStadium API requires a VPN running locally to access your cluster. You will need to install and configure a VPN on the same machine as runner, following the official MacStadium [guide](https://orkadocs.macstadium.com/docs/prerequisites#set-up-vpn).

# Image

The default MacStadium images do not have a git client installed. You will need to create a virtual machine and install the git client and any other software required for your pipelines. Save the virtual machine as a custom image. _You will need to provide the runner with the name of this image as part of the install process._

# Token

The runner uses the MacStadium API to provision virtual machines. Use the Orka command line interface to login to your MacStadium account. The Orka CLI stores your API token in a configuration file in your home directory. _You will need to provide the runner with the token as part of the install process._

* View the contents of the Okra configuration file:
    ```
    $ cat ~/.config/configstore/orka-cli.json
    ```

* Extract the token from the json configuration file:
    ```
    {
        "api-url": "http://10.221.188.100",
        "api-version": "1.1.0",
        token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.Et9HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf58U",
    }
    ```

# Configuration

The runner is configured using environment variables. This article references the below configuration options. See [Configuration]({{< relref "configuration/reference" >}}) for a complete list of configuration options.

* __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.

* __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.

* __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

* __DRONE_ORKA_ENDPOINT__
  : provides your MacStadium API endpoint. The default value is `http://10.221.188.100`.

* __DRONE_ORKA_TOKEN__
  : provides your MacStadium API token.

* __DRONE_VM_CPU__
  : provides the default number of virtual cpus allocated to a virtual machine. The default value is 12 virtual cpus.

* __DRONE_VM_IMAGE__
  : provides the default virtual machine image.

* __DRONE_VM_USERNAME__
  : provides the virtual machine image username. The default value is `admin`.

* __DRONE_VM_PASSWORD__
  : provides the virtual machine image password. The default value is `admin`.

# Installation

The below command creates a container and starts the runner. _Remember to replace the environment variables below with your Drone server details._

```
$ docker run -d \
  -e DRONE_RPC_PROTO=https \
  -e DRONE_RPC_HOST=drone.company.com \
  -e DRONE_RPC_SECRET=super-duper-secret \
  -e DRONE_ORKA_TOKEN=... \
  -e DRONE_VM_IMAGE=... \
  -p 3000:3000 \
  --restart always \
  --name runner \
  drone/drone-runner-macstadium
```

# Verification

Use the `docker logs` command to view the logs and verify the runner successfully established a connection with the Drone server.

```
$ docker logs runner

INFO[0000] starting the server
INFO[0000] successfully pinged the remote server 
```
