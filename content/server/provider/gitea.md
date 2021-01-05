---
date: 2000-01-01T00:00:00+00:00
title: Gitea
author: bradrydzewski
weight: 4
steps: true
toc: true
aliases:
- /installation/providers/gitea
- /install-for-gitea
- /installation/gitea/kubernetes/
- /intro/gitea/
- /installation/gitea
- /installation/gitea/singe-machine
- /installation/gitea/multi-machine
- /installation/gitea/kubernetes
- /intro/gitea/single-machine/
- /intro/gitea/multi-machine/

description: |
  Installation and configure for use with Gitea
---

This article explains how to install the Drone server for Gitea. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

<div class="alert alert-danger">
Please note we strongly recommend installing Drone on a dedicated instance. We do not recommend installing Drone and Gitea on the same instance using docker-compose due to network complications.
</div>

# Preparation

## Create an OAuth Application

Create a Gitea OAuth application. The Consumer Key and Consumer Secret are used to authorize access to Gitea resources.

<div class="alert alert-warn">
The authorization callback URL must match the below format and path, and must use your exact server scheme and host.
</div>

![Application Create](/screenshots/gitea_application_create.png)
![Application View](/screenshots/gitea_application_created.png)

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
$ docker pull drone/drone:1
```

# Configuration

The Drone server is configured using environment variables. This article references a subset of configuration options, defined below. See [Configuration]({{< relref "../reference/_index.md" >}}) for a complete list of configuration options.

* __DRONE_GITEA_CLIENT_ID__
  : Required string value provides your Gitea oauth Client ID.
* __DRONE_GITEA_CLIENT_SECRET__
  : Required string value provides your Gitea oauth Client Secret.
* __DRONE_GITEA_SERVER__
  : Require string value provides your Gitea server address. For example `https://gitea.company.com`, note the `http(s)` otherwise you'll see an error with "unsupported protocol scheme" from Gitea.
* __DRONE_GIT_ALWAYS_AUTH__
  : Optional boolean value configures Drone to authenticate when cloning public repositories.
* __DRONE_RPC_SECRET__
  : Required string value provides the shared secret generated in the previous step. This is used to authenticate the rpc connection between the server and runners. The server and runner must be provided the same secret value.
* __DRONE_SERVER_HOST__
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port. For example `drone.company.com`.
* __DRONE_SERVER_PROTO__
  : Required string value provides your external protocol scheme. This value should be set to http or https. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. For a full list of configuration parameters, please see the configuration reference.

{{< highlight handlebars "linenos=table" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITEA_SERVER={{DRONE_GITEA_SERVER}} \
  --env=DRONE_GITEA_CLIENT_ID={{DRONE_GITEA_CLIENT_ID}} \
  --env=DRONE_GITEA_CLIENT_SECRET={{DRONE_GITEA_CLIENT_SECRET}} \
  --env=DRONE_RPC_SECRET={{DRONE_RPC_SECRET}} \
  --env=DRONE_SERVER_HOST={{DRONE_SERVER_HOST}} \
  --env=DRONE_SERVER_PROTO={{DRONE_SERVER_PROTO}} \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:1
{{< / highlight >}}

# Install Runners

Once your server is up and running you will need to install runners to execute your build pipelines. See our runner installation documentation for detailed installation instructions. 

{{< link "/runner/overview" "Install Runners" >}}
