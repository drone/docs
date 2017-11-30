
+++
date = "2017-07-23T14:39:04+02:00"
title = "Gitea"
url = "install-for-gitea"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-gitea"
  weight = 3
+++

Drone comes with built-in support for the latest stable version of Gitea. To enable Gitea you should configure the Drone container using the following environment variables:

```diff
version: '2'

services:
  drone-server:
    image: drone/agent:{{% version %}}
    ports:
      - 80:8000
      - 9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_GITEA=true
+     - DRONE_GITEA_URL=http://gitea.mycompany.com
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
      - DRONE_SERVER=drone-server:9000
      - DRONE_SECRET=${DRONE_SECRET}
```

# Authentication

Drone will prompt you for a username and password to authenticate. You should use your Gitea username and password. This is unfortunately required due to Gitea lack of oauth2 support.

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.

DRONE_GITEA=true
: Set to true to enable the Gitea driver.

DRONE_GITEA_URL
: Gitea server address.

DRONE_GITEA_GIT_USERNAME
: Optional. Use a single machine account username to clone all repositories.

DRONE_GITEA_GIT_PASSWORD
: Optional. Use a single machine account password to clone all repositories.

DRONE_GITEA_PRIVATE_MODE=false
: Set to true if Gitea is running in private mode.

DRONE_GITEA_SKIP_VERIFY=false
: Set to true to disable SSL verification.
