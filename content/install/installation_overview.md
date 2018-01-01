+++
date = "2017-11-03T15:01:25+02:00"
title = "Installation Overview"
url = "installation"

[menu.install]
  weight = 1
  identifier = "installation"
  parent = "install_overview"
+++

Drone is a lightweight, powerful continuous delivery platform built for containers. Drone is packaged and distributed as a Docker image and can be downloaded from Dockerhub.

```text
docker pull drone/drone:{{% version %}}
```

# Docker Compose

This section provides basic instructions for installing Drone using [docker-compose](https://docs.docker.com/compose/). The below configuration (save it as `docker-compose.yaml`) can be used to start the Drone server with a single agent. It relies on a number of environment variables that you must set before running `docker-compose up`. The variables are described below.

```yaml
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
      - 9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}

  drone-agent:
    image: drone/agent:{{% version %}}
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

Drone integrates with multiple version control providers, configured using environment variables. This example demonstrates basic GitHub integration.

{{% alert %}}
You must register Drone with GitHub to obtain the client and secret. The authorization callback url must match `<scheme>://<host>/authorize`
{{% /alert %}}

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_ORGS=dolores,dogpatch
      - DRONE_ADMIN=johnsmith,janedoe
+     - DRONE_GITHUB=true
+     - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
+     - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

Drone mounts a volume on the host machine to persist the sqlite database.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
      - 9000
+   volumes:
+     - ./drone:/var/lib/drone/
    restart: always
```

Drone needs to know its own address. You must therefore provide the address in `<scheme>://<hostname>` format. Please omit trailing slashes.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
+     - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

Drone agents require access to the host machine Docker daemon.

```diff
services:
  drone-agent:
    image: drone/agent:{{% version %}}
    command: agent
    restart: always
    depends_on: [ drone-server ]
+   volumes:
+     - /var/run/docker.sock:/var/run/docker.sock
```

Drone agents require the server address for agent-to-server communication.

```diff
services:
  drone-agent:
    image: drone/agent:{{% version %}}
    command: agent
    restart: always
    depends_on: [ drone-server ]
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
+     - DRONE_SERVER=drone-server:9000
      - DRONE_SECRET=${DRONE_SECRET}
```

Drone server and agents use a shared secret to authenticate communication. This should be a random string of your choosing and should be kept private.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
+     - DRONE_SECRET=${DRONE_SECRET}
  drone-agent:
    image: drone/agent:{{% version %}}
    environment:
      - DRONE_SERVER=drone-server:9000
      - DRONE_DEBUG=true
+     - DRONE_SECRET=${DRONE_SECRET}
```

Drone registration is closed by default. This example enables open registration for users that are members of approved GitHub organizations.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_OPEN=true
+     - DRONE_ORGS=dolores,dogpatch
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

Drone administrators should also be enumerated in your configuration.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_ORGS=dolores,dogpatch
+     - DRONE_ADMIN=johnsmith,janedoe
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```
