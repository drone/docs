+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Services"
weight = 20
toc = true


[menu.main]
	parent="usage"
+++

# Overview

Drone provides a `services` section in the Yaml file used for defining service containers, such as databases. Service containers share the same `localhost` network as your build containers.

Example configuration that composes a Postgres and Redis container:

```yaml
services:
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
image: postgres
image: postgres:9.2
image: library/postgres:9.2
image: index.docker.io/library/postgres:9.2
```

Use the `pull` attribute to instruct Drone to always pull the latest Docker image. This helps ensure you are always testing your code against the latest image:

```yaml
services:
  database:
    image: postgres
    pull: true
```

# Authorization

Drone supports private images that require password authentication. You can use the command line utility to register authentication credentials:

```
drone secrets add octocat/hello-world REGISTRY_USERNAME octocat
drone secrets add octocat/hello-world REGISTRY_PASSWORD pa55word
drone secrets add octocat/hello-world REGISTRY_EMAIL octocat@github.com
```

You can alternatively specify the authentication credentials directly in the Yaml:

```yaml
services:
  database:
    image: registry.yourcompany.com/your/image:tag
    auth_config:
      username: octocat
      password: pa55word
      email: octocat@github.com
```

# Networking

Service containers are available at the `localhost` or `127.0.0.1` address.

Drone deviates from the default Docker compose networking model to mirror a traditional development environment, where services are typically accessed at `localhost` or `127.0.0.1`. To achieve this, we create a per-build network where all containers share the same network and IP address.

# Linking

Please see the above networking section -- because your build and service containers share the same network linking is not possible. Instead you may access your services at `127.0.0.1` similar to a traditional, local development environment.

# Environment

Use the environment section to pass environment variables to your service containers. Many database images, including the official postgres and mysql images, use environment variables to configure the database on startup.

```yaml
services:
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=mysecretpassword
```

# Volumes

Use the `volumes` attribute to mount folders on your host machine into your service container. These are [Docker volumes](https://docs.docker.com/engine/userguide/dockervolumes/) and therefore use the same `<host>:<container>` declaration conventions:

```yaml
services:
  database:
    image: postgres
    volumes:
      - /some/path/on/host:/path/in/container
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Privileged

Use the privileged attribute to run your service in a privileged Docker container:

```yaml
services:
  dind:
    image: docker:dind
    privileged: true
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.
