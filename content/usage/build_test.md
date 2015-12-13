+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Build & Test"
weight = 7
menu = "usage"
toc = true
+++

# Overview

# Images

Drone supports any valid Docker image from any Docker registry for your build environment. Declare your image using standard image names in short form, or fully qualified:

```
image: golang
image: golang:1.9
image: library/golang
image: library/golang:1.9
image: index.docker.io/library/golang:1.9
```

Provide your registry credentials if your build image is private:

```
build:
  image: foo/bar
  auth_config:
    username: octocat
    password: password
```


# Environment

Use the environment section to inject environment variables into your build environment:

```
build:
  image: golang:1.5
  environment:
    - GO15VENDOREXPERIMENT=0
    - GOOS=linux
    - GOARCH=amd64
    - CGO_ENABLED=0
```

Variable expansion is not supported. The following example __will not work__:

```
environment:
  - PATH=$PATH:/go
```

# Commands

Build commands execute sequentially in a shell environment with the `-x` flag. The `-x` flag instructs the shell environment to exit immediately if a command returns a non-zero exit code.

```
build:
  image: golang:1.5
  commands:
    - go get
    - go build
    - go test
```

# Volumes

Use the `volumes` attribute to mount folders on your host machine into your build container. These are Docker volumes and therefore use the same `<host>:<container>` declaration conventions:

```
build:
  image: golang
  volumes:
    - /some/path/on/host:/path/in/container
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Privileged

Use the `privileged` attribute to run your build in a privileged Docker container:

```
build:
  image: golang
  privileged: true
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Timeouts

There is a default 60 minute timeout for your build. This can be increased by a Drone administrator from your repository settings screen.

# Skip Commits

Instruct Drone to skip builds by adding `[ci skip]` to your commit message.

# Skip Branches

Instruct Drone to skip branches by including a branch white-list in your `.drone.yml`

```
branches:
  - master
  - develop
```

# Known Issues

Drone expects the container to execute with `UID 0` (root) inside the container. If your image defines an alternate `USER` your build will fail with the following error message:

```
/bin/sh: 49: cannot create /etc/apt/apt.conf.d/90forceyes: Permission denied
```
