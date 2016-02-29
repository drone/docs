+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Service Containers"
weight = 6
menu = "usage"
toc = true
+++

# Overview

Drone uses the `compose` section of the `.drone.yml` to specify supporting containers (ie service containers) that should be started and linked to your build container. The `compose` section of the `.drone.yml` is modeled after the `docker-compose.yml`:

```yaml
---
compose:
  [container_name:]
    image: [image_name]
    [options]
```

Example configuration that composes a Postgres and Redis container:

```yaml
---
compose:
  cache:
    image: redis
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=mysecretpassword
```

# Images

Drone supports any valid Docker image from any Docker registry:

```yaml
---
image: postgres
image: postgres:9.2
image: library/postgres:9.2
image: index.docker.io/library/postgres:9.2
```

Provide your registry credentials if your service container image is private:

```yaml
---
compose:
  database:
    image: myrepo/customdb:latest
    auth_config:
      username: octocat
      password: password
```

Use the `pull` attribute to instruct Drone to always pull the latest Docker image. This helps ensure you are always testing your code against the latest image:

```yaml
---
compose:
  database:
    image: postgres
    pull: true
```

# Workspace

The build workspace is automatically mounted into each service container in the ```/drone/src/<repository>``` directory inside the ```/drone``` volume.
If needed, this provides access to the source code and build artifacts.

Here is an example when using Github:

      /drone/src/github.com/octocat/hello-world

See also: [Build & Test](../build_test/#workspace:fb92aa3346185c57f15afda861d465a3)

# Networking

Service containers are available at the `localhost` or `127.0.0.1` address.

Drone deviates from the default Docker compose networking model to mirror a traditional development environment, where services are typically accessed at `localhost` or `127.0.0.1`. To achieve this, we create a per-build network where all containers share the same network and IP address.

# Linking

Please see the above networking section -- because your build and service containers share the same network linking is not possible. Instead you may access your services at `127.0.0.1` similar to a traditional, local development environment.

# Environment

Use the environment section to pass environment variables to your service containers. Many database images, including the official postgres and mysql images, use environment variables to configure the database on startup.

```yaml
---
compose:
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=mysecretpassword
```

# Volumes

Use the `volumes` attribute to mount folders on your host machine into your service container. These are [Docker volumes](https://docs.docker.com/engine/userguide/dockervolumes/) and therefore use the same `<host>:<container>` declaration conventions:

```yaml
---
compose:
  database:
    image: postgres
    volumes:
      - /some/path/on/host:/path/in/container
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Devices

Use the `devices` attribute to map devices from your host machine into your service container. These are [Docker devices](https://docs.docker.com/compose/compose-file/#devices) and therefore use the same declaration conventions:

```yaml
---
compose:
  database:
    image: postgres
    devices:
      - "/dev/ttyUSB0:/dev/ttyUSB0"
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Privileged Mode

Use the privileged attribute to run your service in a privileged Docker container:

```yaml
---
compose:
  dind:
    image: docker:dind
    privileged: true
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Docker Compose

You should think of the `.drone.yml` as a competing format to the `docker-compose.yml` that specializes in testing and build pipelines. We encourage you to embrace the `.drone.yml` for testing instead of trying to hack together a homegrown solution with a `docker-compose.yml`.

Using a `docker-compose.yml` is possible but will require direct access to a Docker daemon. Drone does not expose a Docker daemon to your build by default for security reasons. You can provide access to a Docker daemon by mounting the host machines Docker socket as a [volume](#volumes:bfc9941b6b6fd7b4ef09dd0ccd08af0c) or by running [Docker in Docker](#docker-in-docker:bfc9941b6b6fd7b4ef09dd0ccd08af0c) as a service container.

# Docker in Docker

TODO

# Examples

TODO
