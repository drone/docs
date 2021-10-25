---
date: 2000-01-01T00:00:00+00:00
title: GitLab
author: bradrydzewski
steps: true
toc: true
weight: 4
aliases:
- /installation/gitlab
- /installation/providers/gitlab
- /install-for-gitlab
- /admin/setup-gitlab/
- /intro/gitlab/
- /intro/gitlab/single-machine/
- /intro/gitlab/multi-machine/
- /intro/gitlab/kubernetes
- /installation/gitlab/single-machine
description: |
  Installation and configure for use with GitLab
---

This article explains how to install the Drone server for GitLab. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

# Preparation

## Create an OAuth Application

Create a GitLab OAuth application. The Consumer Key and Consumer Secret are used to authorize access to GitLab resources.

<div class="alert alert-warn">
The authorization callback URL must match the below format and path, and must use your exact server scheme and host.
</div>

![Application Create](/screenshots/gitlab_token_create.png)
![Application View](/screenshots/gitlab_token_created.png)

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

* __DRONE_GITLAB_CLIENT_ID__
  : Required string value provides your GitLab oauth Client ID.
* __DRONE_GITLAB_CLIENT_SECRET__
  : Required string value provides your GitLab oauth Client Secret.
* __DRONE_GITLAB_SERVER__
  : Option string value provides your GitLab server url. The default value is the gitlab.com server address at `https://gitlab.com`.
* __DRONE_GIT_ALWAYS_AUTH__
  : Optional boolean value configures Drone to authenticate when cloning public repositories. This should only be enabled when using self-hosted GitLab with private mode enable.
* __DRONE_RPC_SECRET__
  : Required string value provides the shared secret generated in the previous step. This is used to authenticate the rpc connection between the server and runners. The server and runner must be provided the same secret value.
* __DRONE_SERVER_HOST__
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port. For example, `drone.domain.com`
* __DRONE_SERVER_PROTO__
  : Required string value provides your external protocol scheme. This value should be set to `http` or `https`. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. _Remember to replace the placeholder values below with the appropriate values._

{{< highlight bash "linenos=table,hl_lines=3-8" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITLAB_SERVER=https://gitlab.com \
  --env=DRONE_GITLAB_CLIENT_ID=05136e57d80189bef462 \
  --env=DRONE_GITLAB_CLIENT_SECRET=7c229228a77d2cbddaa61ddc78d45e \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_SERVER_HOST=drone.company.com \
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
