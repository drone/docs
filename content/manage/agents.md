+++
date = "2015-12-05T16:00:21-08:00"
title = "Agents"
weight = 3
menu = "installation"
toc = true
+++

# Overview

Drone agents poll the Drone server queue for pending builds. You can install one or many Drone agents on one or many servers to scale your build infrastructure. You can add and remove agents as needed, without restarting the Drone server, for dynamic scaling.

# Authorization

Drone authorizes agents using a special access token. The token is currently written to your drone logs when you first start the server. You will need to find the secret in your server logs before we proceed to the next step.

```
msg="agents can connect with token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
```

# Installation

Get started by downloading the Docker image from:

```
docker pull drone/drone:0.5.0
```

Create a `/etc/dronerc` file to hold your configuration parameters. The only two required parameters are the server address and server token. Please see the prior section of this document for instructions retrieving your server token.

```
DRONE_SERVER=http://drone.server.com
DRONE_TOKEN=...
```

Create and run your container with Docker mounted as a volume:

```
docker run \
	--volume /var/run/docker.sock:/var/run/docker.sock \
	--env-file /etc/dronerc \
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
`DRONE_TOKEN`               | drone server access token
`DRONE_DEBUG`               | drone started in debug mode
`DRONE_BACKOFF`             | drone server backoff duration
`DRONE_PLUGIN_PULL`         | plugin updates are pull automatically
`DRONE_PLUGIN_NAMESPACE`    | plugin namespace in dockerhub, defaults to `plugins`
`DRONE_PLUGIN_WHITELIST`    | plugin whitelist, defaults to `plugins/*`
`DRONE_PLUGIN_PRIVILEGED`   | plugins automatically granted privileged mode
`DRONE_PLUGIN_NETRC`        | plugins automatically granted netrc access
`DOCKER_HOST`               | docker host address
`DOCKER_TLS_VERIFY`         | docker requires tls verification
`DOCKER_CERT_PATH`          | docker certificate path
`DOCKER_MAX_PROCS`          | docker concurrent build processes
`DOCKER_OS`                 | docker operating system (linux)
`DOKCER_ARCH`               | docker architecture (amd64)
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

# Troubleshooting

Please use the Drone logs troubleshoot issues:

```
docker logs drone-agent
```

For more verbose logging enable debugging in your `dronerc` file:

```
DRONE_DEBUG=true
```
