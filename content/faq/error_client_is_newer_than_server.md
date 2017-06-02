---
date: 2017-04-15T14:39:04+02:00
title: "Error: client is newer than server"
url: error-client-is-newer-than-server
---

If you are running a legacy version of the Docker daemon you will encounter the following error message:

```nohighlight
ERROR: Error response from daemon: client is newer than
server (client API version: 1.26, server API version: 1.24)
```

This can be resolved by configuring your agent with the following environment variable:

```nohighlight
DOCKER_API_VERSION=1.24
```

This error message is encountered because Docker has a versioned protocol, and legacy versions of Docker prevent newer versions of the protocol from connecting. The solution is to downgrade the protocol using the DOCKER_API_VERSION environment variable.
