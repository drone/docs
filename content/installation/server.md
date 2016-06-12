+++
date = "2015-12-05T16:00:21-08:00"
title = "Server"
weight = 1
toc = true

[menu.main]
	parent="installation"
+++

# Overview

Drone ships as a single binary file inside a minimalist 20 MB Docker image. Docker is the only dependency. We recommend using Drone with the latest stable version of Docker and officially support Docker 1.N-1 (latest minor release minus one).

# Installation

Get started by downloading the Docker image:

```
docker pull drone/drone:0.5
```

Configuration parameters are provided to Drone as environment variables. Please read the documentation to configure your service integrations (ie GitHub) before completing this step.

```
DRONE_GITHUB_CLIENT=...
DRONE_GITHUB_SECRET=...
```

Create a secret passcode that will be used to authenticate your builds agents. We recommend using a [random password generator](http://correcthorsebatterystaple.net/) to create this secret passcode.

```
DRONE_SECRET=...
```

Choose a strategy for authorizing and registering users. See the user section of the documentation for more details. This example enables open registration and grants administrative access:

```
DRONE_OPEN=true
DRONE_ADMIN=bradrydzewski,octocat
```

Create and run your container:

```
docker run \
  --env DRONE_GITHUB_CLIENT=... \
  --env DRONE_GITHUB_SECRET=... \
  --env DRONE_SECRET=... \
  --env DRONE_OPEN=true  \
  --env DRONE_ADMIN=...  \
  --volume /var/lib/drone:/var/lib/drone \
  --restart=always \
  --publish=80:8000 \
  --detach=true \
  --name=drone \
  drone/drone:0.5
```

Note the above example mounts a volume on the host machine. The default configuration uses a sqlite database and should therefore be mounted on the host machine as a volume to avoid data loss.

```
--volume /var/lib/drone:/var/lib/drone
```

# Configuration

For a full list of configuration options, see the [Drone Server Configuration Reference](../../reference/configuration/server).

# Logging

Please use the Drone logs troubleshoot issues:

```
docker logs drone
```

For more verbose logging run drone with the following environment variable:

```
DRONE_DEBUG=true
```
