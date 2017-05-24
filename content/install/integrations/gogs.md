+++
date = "2017-04-15T14:39:04+02:00"
title = "Gogs"
url = "install-for-gogs"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-gogs"
  weight = 3
+++

Drone comes with built-in support for the latest stable version of Gogs. To enable Gogs you should configure the Drone container using the following environment variables:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_GOGS=true
+     - DRONE_GOGS_URL=http://gogs.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}

  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

# Authentication

Drone will prompt you for a username and password to authenticate. You should use your Gogs username and password. This is unfortunately required due to Gogs lack of oauth2 support.

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.

DRONE_GOGS=true
: Set to true to enable the Gogs driver.

DRONE_GOGS_URL
: Gogs server address.

DRONE_GOGS_GIT_USERNAME
: Optional. Use a single machine account username to clone all repositories.

DRONE_GOGS_GIT_PASSWORD
: Optional. Use a single machine account password to clone all repositories.

DRONE_GOGS_PRIVATE_MODE=false
: Set to true if Gogs is running in private mode.

DRONE_GOGS_SKIP_VERIFY=false
: Set to true to disable SSL verification.
