---
date: 2000-01-01T00:00:00+00:00
title: Bitbucket Cloud
author: bradrydzewski
steps: true
toc: true
weight: 5
aliases:
- /installation/providers/bitbucket-cloud
- /install-for-bitbucket-cloud
- /installation/bitbucket-cloud
- /installation/bitbucket-cloud/single-machine/
- /installation/bitbucket-cloud/multi-machine/
- /installation/bitbucket-cloud/kubernetes
description: |
  Installation and configure for use with Bitbucket Cloud
---

This article explains how to install the Drone server for Bitbucket Cloud. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

# Preparation

## Provision an Instance

The Drone server should be installed on a server or virtual machine (using your cloud provider of choice) with standard http and https ports open. The instance must be publicly accessible by domain name or IP address to receive inbound webhooks from Bitbucket.

_When installing the Drone server on your laptop for testing purposes, we recommend using a service like [ngrok](https://ngrok.com/) to provide your Drone server with a publicly accessible domain name._

## Create an OAuth Application

Create a Bitbucket OAuth application. The Consumer Key and Consumer Secret are used to authorize access to Bitbucket resources.

<div class="alert alert-warn">
The authorization callback URL must match the below format and path, and must use your exact server scheme and host.
</div>

![Application Create](/screenshots/bitbucket_application_create.png)
![Application View](/screenshots/bitbucket_application_list.png)

## Create a Shared Secret
Create a shared secret to authenticate communication between runners and your central Drone server.

You can use openssl to generate a shared secret:

```
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

# Download

The Drone server is distributed as a lightweight Docker image. The image is self-contained and does not have any external dependencies.

```
$ docker pull drone/drone:2
```

# Configuration

The Drone server is configured using environment variables. This article references a subset of configuration options, defined below. See [Configuration]({{< relref "../reference/_index.md" >}}) for a complete list of configuration options.

* __DRONE_BITBUCKET_CLIENT_ID__
  : Required string value provides your Bitbucket oauth Client ID.
* __DRONE_BITBUCKET_CLIENT_SECRET__
  : Required string value provides your Bitbucket oauth Client Secret.
* __DRONE_RPC_SECRET__
  : Required string value provides the shared secret generated in the previous step. This is used to authenticate the rpc connection between the server and runners. The server and runner must be provided the same secret value.
* __DRONE_SERVER_HOST__
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port. For example `drone.company.com`.
* __DRONE_SERVER_PROTO__
  : Required string value provides your external protocol scheme. This value should be set to http or https. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. _Remember to replace the placeholder values below with the appropriate values._

{{< highlight handlebars "linenos=table" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_BITBUCKET_CLIENT_ID={{DRONE_BITBUCKET_CLIENT_ID}} \
  --env=DRONE_BITBUCKET_CLIENT_SECRET={{DRONE_BITBUCKET_CLIENT_SECRET}} \
  --env=DRONE_RPC_SECRET={{DRONE_RPC_SECRET}} \
  --env=DRONE_SERVER_HOST={{DRONE_SERVER_HOST}} \
  --env=DRONE_SERVER_PROTO={{DRONE_SERVER_PROTO}} \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:2
{{< / highlight >}}

# Install Runners

Once your server is up and running you will need to install runners to execute your build pipelines. See our runner installation documentation for detailed installation instructions. 

{{< link "/runner/overview" "Install Runners" >}}
