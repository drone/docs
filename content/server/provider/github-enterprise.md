---
date: 2000-01-01T00:00:00+00:00
title: GitHub Enterprise
author: bradrydzewski
weight: 1
steps: true
toc: true

related:
  items:
  - name: Setup TLS
    path: ../https.md
  - name: Setup Prometheus Metrics
    path: ../metrics.md
  - name: Setup Administrator Accounts
    path: ../user/admin.md
  - name: Database Options
    path: ../storage/database.md

description: |
  Installation and configure for use with GitHub Enterprise
---

This article explains how to install the Drone server for GitHub Enterprise. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

# Preparation

## Create an OAuth Application

Create a GitHub OAuth application. The Consumer Key and Consumer Secret are used to authorize access to GitHub resources.

<div class="alert alert-warn">
The authorization callback URL must match the below format and path, and must use your exact server scheme and host.
</div>

![Application Create](/screenshots/github_application_create.png)
![Application View](/screenshots/github_application_created.png)

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

* __DRONE_GITHUB_CLIENT_ID__
  : Required string value provides your GitHub oauth Client ID.
* __DRONE_GITHUB_CLIENT_SECRET__
  : Required string value provides your GitHub oauth Client Secret.
* __DRONE_GITHUB_SERVER__
  : Required url value provides your GitHub Enterprise server address. For example `https://github.commpany.com`.
* __DRONE_GIT_ALWAYS_AUTH__
  : Optional boolean value configures Drone to authenticate when cloning public repositories. This should only be enabled when using GitHub Enterprise with private mode enable.
* __DRONE_RPC_SECRET__
  : Required string value provides the shared secret generated in the previous step. This is used to authenticate the rpc connection between the server and runners. The server and runner must be provided the same secret value.
* __DRONE_SERVER_HOST__
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port. For example `drone.company.com`.
* __DRONE_SERVER_PROTO__
  : Required string value provides your external protocol scheme. This value should be set to http or https. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. For a full list of configuration parameters, please see the configuration reference.

```handlebars {linenos=table,linenostart=1}
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITHUB_SERVER={{DRONE_GITHUB_SERVER}} \
  --env=DRONE_GITHUB_CLIENT_ID={{DRONE_GITHUB_CLIENT_ID}} \
  --env=DRONE_GITHUB_CLIENT_SECRET={{DRONE_GITHUB_CLIENT_SECRET}} \
  --env=DRONE_RPC_SECRET={{DRONE_RPC_SECRET}} \
  --env=DRONE_SERVER_HOST={{DRONE_SERVER_HOST}} \
  --env=DRONE_SERVER_PROTO={{DRONE_SERVER_PROTO}} \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:2
```

# Install Runners

Once your server is up and running you will need to install runners to execute your build pipelines. See our runner installation documentation for detailed installation instructions. 

{{< link "/runner/overview" "Install Runners" >}}
