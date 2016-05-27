+++
date = "2015-12-05T16:00:21-08:00"
title = "drone exec"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++


# Overview

The exec subcommand gives you the ability to run your build locally, using your local Docker daemon. This is useful for testing and debugging your code locally, using an environment that __closely__ matches the Drone server environment.

Example command to run your build locally:

```
$ drone exec
```

# Docker Setup

The exec command requires you to have a Docker instance running locally. Please be sure the standard Docker environment variables are configured so the command line utility can connect to the Daemon:

```
DOCKER_HOST=...
DOCKER_CERT_PATH=...
DOCKER_TLS_VERIFY=...
```

# Build Params

The full set of build parameters, typically sourced from your post-commit hook, can be set from the command line. Please use `drone help exec` for the complete list of available parameters.

```
$ drone exec \
    --repo.owner=octocat \
    --repo.name=hello-world \
    --build.number=1 \
    --build.event=push
```

# Secrets

Specify secrets for your local build:

```
$ drone exec \
    --secret DOCKER_USERNAME=... \
    --secret DOCKER_PASSWORD=...
```

# Matrix

Specify matrix values for your local build:

```
$ drone exec \
    --matrix REDIS_VERSION=3.2 \
    --matrix GOLANG_VERSION=1.6.2
```

# Workspace

The command line utility mounts your local working directory into the build containers using Docker volumes. If your source code needs to be mounted into a specific directory please declare your workspace in the Yaml:

```
workspace:
  base: /go
  path: src/github.com/octocat/hello-world
```
