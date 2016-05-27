+++
date = "2015-12-05T16:00:21-08:00"
title = "Agents"
weight = 2
toc = true

[menu.main]
	parent="installation"
+++

# Overview

Drone agents poll the Drone server queue for pending builds. You can install one or many Drone agents on one or many servers to scale your build infrastructure. You can add and remove agents as needed, without restarting the Drone server, for dynamic scaling.

# Installation

Get started by downloading the Docker image from:

```
docker pull drone/agent:0.5
```

The only two required parameters are the server address and secret token. We will provide these parameters to Docker as environment variables using the `--env` flag. Please see the prior section of this document for instructions retrieving your secret token.

```
DRONE_SERVER=http://drone.server.com
DRONE_SECRET=...
```

Create and run your container with Docker mounted as a volume:

```
docker run \
  --env DRONE_SERVER=... \
  --env DRONE_SECRET=... \
	--volume /var/run/docker.sock:/var/run/docker.sock \
	--restart=always \
	--detach=true \
	--name=drone-agent \
	drone/drone:0.5 agent
```


# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration value that should work for the majority of installations.

NAME                        | DESC
----------------------------|--------------------------------------------------------
`DRONE_SERVER`              | drone server address
`DRONE_SECRET`              | drone server secret token
`DRONE_DEBUG`               | drone started in debug mode
`DRONE_PING`                | drone server ping frequency
`DRONE_BACKOFF`             | drone server backoff duration
`DRONE_TIMEOUT`             | build timeout for console inactivity
`DRONE_MAX_LOGS`            | build log size limit, defaults to 5mb
`DRONE_PLUGIN_PULL`         | plugin updates are pull automatically
`DRONE_PLUGIN_NAMESPACE`    | plugin namespace in dockerhub, defaults to `plugins`
`DRONE_PLUGIN_PRIVILEGED`   | plugins automatically granted privileged mode
`DOCKER_HOST`               | docker host address
`DOCKER_TLS_VERIFY`         | docker requires tls verification
`DOCKER_CERT_PATH`          | docker certificate path
`DOCKER_MAX_PROCS`          | docker concurrent build processes
`DOCKER_OS`                 | docker operating system (linux)
`DOCKER_ARCH`               | docker architecture (amd64)
`HTTP_PROXY`                | http proxy server
`HTTPS_PROXY`               | https proxy server
`NO_PROXY`                  | proxy server exceptions

# Connectivity

Agents automatically connect to the central server and begin polling for pending builds. If the server connection is interrupted the agents will use exponential backoff to automatically reconnect:

```
WARN[0000] Attempting to reconnect in 15s
WARN[0000] Attempting to reconnect in 15s
```

If the server connection is interrupted existing live-streaming logs may not resume. Agents will use exponential backoff to send build updates and completed logs when the connection resumes. You can therefore, in most cases, safely restart the Drone server with limited impact to your users.

# Logging

Please use the Drone logs troubleshoot issues:

```
docker logs drone-agent
```

For more verbose logging enable debugging:

```
DRONE_DEBUG=true
```
