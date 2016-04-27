+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "FAQ"
weight = 33
menu = "usage"
toc = true
+++

# Why does my build fail?

There is no easy way to answer this question. There are a number of reasons your build could pass locally and fail when running in Drone. Most people _incorrectly_ assume there must be a bug with Drone. The most common reason for build failure is incorrect configuration.

We recommend using the Drone command line utility for local testing and debugging. If you are unable to resolve the issue locally you can use the community [support forums](https://gitter.im/drone/drone) to ask for help. If you engage the community for help please provide sample Yaml configurations and build output.

# How can I trigger builds?

Builds are automatically triggered by post-commit hooks from your version control system. You can use the API or command line tools to re-start or re-run existing builds with an auto-incremented build number.

# How can I schedule builds?

There are no built-in scheduling capabilities (nor or any planned). External scheduling tools such as cron can be used in conjunction with the official API or command line tools to execute nightly builds.

# How do I test Docker images?

You can use the Docker plugin to build and publish images. If you need to build and _test_ your image you can mount the Docker socket into your build environment to work directly with the host machines Docker daemon.

```
script:
  build:
    image: docker
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    commands:
      - docker build ...
      - docker run ...
      - docker push ...
```

# How do I use Docker compose?

The `service` section of the `.drone.yml` is the official approach for composing service containers. You can alternatively mount the Docker socket into your build environment to work directly with the host machines Docker daemon and Docker compose.

```
script:
  build:
    image: ...
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    commands:
      - docker-compose up
      - ...
```

# Cannot connect to a database.

The most common issue is that databases take time to start and accept connections. We recommend adding `sleep` to your build script to give the database enough time to start. You should also make sure you are connecting to your database using `localhost`.

# Cannot create a database.

Most official docker images use environment variables to create databases and users at runtime. Please see the official image documentation for more details.

# Cannot provide a database user or password.

 Most official docker images use environment variables to create the default user and password. This includes the ability to authenticate with no password. Please see the official image documentation for more details.
