---
date: 2000-01-01T00:00:00+00:00
title: Gitee
author: kit101
weight: 4
steps: true
toc: true
aliases:
- /installation/providers/gitee
- /install-for-gitee
- /admin/setup-gitee
- /installation/gitee/
- /installation/gitee/single-machine/
- /installation/gitee/multi-machine/
- /installation/gitee/kubernetes/
description: |
  Installation and configure for use with Gitee
---

This article explains how to install the Drone server for Gitee. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

# Preparation

## Create an OAuth Application

Create a Gitee OAuth application. The Consumer Key and Consumer Secret are used to authorize access to Gitee resources.

<div class="alert alert-warn">
The authorization callback URL must match the below format and path, and must use your exact server scheme and host.
</div>

![Application Create](/screenshots/gitee_token_create.png)
![Application View](/screenshots/gitee_token_created.png)

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

* __DRONE_GITEE_CLIENT_ID__
  : Required string value provides your Gitee oauth Client ID.
* __DRONE_GITEE_CLIENT_SECRET__
  : Required string value provides your Gitee oauth Client Secret.
* __DRONE_GITEE_SERVER__
  : Optional url value provides Gitee server address. The default value is the gitee.com server address at `https://gitee.com`.
* __DRONE_GITEE_API_SERVER__
  : Optional string value provides Gitee api server address. The default value is `https://gitee.com/api/v5`.
* __DRONE_RPC_SECRET__
  : Required string value provides the shared secret generated in the previous step. This is used to authenticate the rpc connection between the server and runners. The server and runner must be provided the same secret value.
* __DRONE_SERVER_HOST__
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port. For example, `drone.domain.com`
* __DRONE_SERVER_PROTO__
  : Required string value provides your external protocol scheme. This value should be set to `http` or `https`. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. _Remember to replace the placeholder values below with the appropriate values._

{{< highlight bash "linenos=table,hl_lines=3-7" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITEE_CLIENT_ID=f7018cdd7c2a2515eb0cc3eeea039a3aeda0991a520c9e6f7eca37b97761de20 \
  --env=DRONE_GITEE_CLIENT_SECRET=a29f1465460f620d1238b6ebf207ae6778a6bfd074b3c7befe72d5f9647ed02c \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_SERVER_HOST=drone.domain.com \
  --env=DRONE_SERVER_PROTO=https \
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
