---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
toc: true
aliases:
- /installation/overview
- /installation
- /administration
- /admin/installation-guide
- /0.5/install/server-configuration
description: |
  Installation overview.
---

This section of the documentation will help you install and configure the Drone _Server_ and one or many _Runners_. A runner is a standalone daemon that polls the server for pending pipelines to execute.

# Server Installation

Drone integrates seamlessly with popular _Source Control Management_ providers. Please choose your provider to get started.

{{< link "/server/provider/github" >}}
{{< link "/server/provider/gitlab" >}}
{{< link "/server/provider/gogs" >}}
{{< link "/server/provider/gitea" >}}
{{< link "/server/provider/bitbucket-cloud" >}}
{{< link "/server/provider/bitbucket-server" >}}

# Server Upgrades

Drone server upgrades are meant to be simple and safe. You can replace the old Docker image with the new Docker image and Drone will handle everything else (automatic database migrations, etc).

If a release includes breaking changes or requires manual upgrade it will be specified in the release notes. _As a general rule we try to maintain strict backward compatibility and avoid breaking changes, even in major releases._

```sh {linenos=table}
# terminate the drone server
docker stop drone
docker rm drone

# pull the latest major release
docker pull drone/drone:2

# re-start the drone server using the
# newer docker image
docker run ... 
```

# Runner Installation

Drone _Runners_ are standalone daemons that poll the server for pending pipelines to execute. There are different types of runners optimized for different use cases and runtime environments. Once the server is successfully installed you must install one or more runners.

{{< link "/runner/overview" "See the Runner Installation Guide" >}}
