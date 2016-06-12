+++
date = "2015-12-05T16:00:21-08:00"
title = "Agent Config"
weight = 2

[menu.main]
	parent="configuration"
+++

Below is a complete reference for all Drone Agent settings. These can all be passed via environment variables, as shown in the *NAME* column.

NAME                        | DESC
----------------------------|--------------------------------------------------------
`DRONE_SERVER`              | Full URL (with protocol) to the Drone Server.
`DRONE_SECRET`              | Shared secret. Must match the value on Drone Server.
`DRONE_DEBUG`               | Set to `true` to see verbose debugging log output.
`DRONE_PING`                | Drone Server ping frequency (in seconds).
`DRONE_BACKOFF`             | Drone Server reconnect backoff duration.
`DRONE_TIMEOUT`             | Build timeout for console inactivity.
`DRONE_MAX_LOGS`            | Build log size limit, defaults to 5mb.
`DRONE_PLUGIN_PULL`         | If `true`, plugin updates are pull automatically.
`DRONE_PLUGIN_NAMESPACE`    | Plugin namespace in dockerhub, defaults to `plugins`.
`DRONE_PLUGIN_PRIVILEGED`   | Space-separated list of plugins automatically granted privileged mode.
`DOCKER_HOST`               | Docker host address.
`DOCKER_TLS_VERIFY`         | Docker requires tls verification.
`DOCKER_CERT_PATH`          | Docker certificate path.
`DOCKER_MAX_PROCS`          | Docker concurrent build processes.
`DOCKER_OS`                 | Docker operating system (linux).
`DOCKER_ARCH`               | Docker architecture (amd64).
`HTTP_PROXY`                | HTTP proxy server.
`HTTPS_PROXY`               | HTTPS proxy server.
`NO_PROXY`                  | Proxy server exceptions.
