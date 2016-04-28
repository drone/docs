+++
date = "2015-12-05T16:00:21-08:00"
draft = true
title = "Steps"
weight = 20
menu = "usage"
toc = true
+++

# Overview

Drone uses the `script` section of the `.drone.yml` to define your Docker build environment and your build and test instructions. The following is an example build definition:

```yaml
---
build:
  image: golang
  commands:
    - go get
    - go build
    - go test
```

This example downloads and runs a `golang` container and executes the specified commands inside the container. Your source code, previously fetched during the `clone` step, is provided to the container as a Docker volume.

# Images

Drone supports any valid Docker image from any Docker registry:

```yaml
---
image: golang
image: golang:1.5
image: library/golang:1.5
image: index.docker.io/library/golang:1.5
```

Use the `pull` attribute to instruct Drone to always pull the latest Docker image. This helps ensure you are always testing your code against the latest image:

```yaml
image: golang
pull: true
```

# Environment

Use the environment section to inject environment variables into your build environment:

```yaml
script:
  build:
    image: golang
    environment:
      - GOOS=linux
      - GOARCH=amd64
      - CGO_ENABLED=0
```

Variable expansion is not supported. The following example __will not work__:

```yaml
environment:
  - PATH=$PATH:/go
```

# Commands

Build commands execute sequentially in a shell environment with the `-e` flag. The `-e` flag instructs the shell environment to exit immediately if a command returns a non-zero exit code.

```yaml
script:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test
```

# Multiple Steps

Build commands can be split into multiple named steps. The below example let's us split our frontend build (node) and our backend build (go). Each step executes in its own isolated container environment while sharing the same build workspace.

```yaml
script:
  frontend:
    image: node:5.3.0
    commands:
      - npm install
      - npm test
  backend:
    image: golang:1.5
    commands:
      - go get
      - go build
      - go test
```

# Filters

Limit when a build step should is executed using filters. The below example splits our build into two steps. The second build step, building the distribution, is only executed when we tag a release. This reduces the build time for `push` and `pull_request` events by removing un-necessary steps:

```yaml
---
build:
  test:
    image: golang:1.5
    commands:
      - go get
      - go test
  bundle:
    image: golang:1.5
    commands:
      - go build
      - tar -cvf bundle.tar myapp
    when:
      event: tag
```

Only execute a step when building a `push` or `tag` event:

```yaml
script:
  build:
    image: golang
    when:
      event: [push, tag]
```

Only executes a step when building the `master` branch:

```yaml
script:
  build:
    image: golang
    when:
      branch: master
```

# Dependencies

Cloning private submodules or private dependencies (using `npm` or `go get`) requires remote authentication. We recommend using `git+https` for submodules or dependencies, allowing Drone to authenticate to your remote repository using a `.netrc` file injected by default into your build environment.

Instead of using the `git+ssh` url in your `package.json` file:

```js
{
    "library": "git+ssh://git@github.com:octocat/library.git"
}
```

Use the `git+https` url in your `package.json` file:

```js
{
    "library": "https://github.com/octocat/library.git"
}
```

# Volumes

Use the `volumes` attribute to mount folders on your host machine into your build container. These are [Docker volumes](https://docs.docker.com/engine/userguide/dockervolumes/) and therefore use the same `<host>:<container>` declaration conventions:

```yaml
---
build:
  image: golang
  volumes:
    - /some/path/on/host:/path/in/container
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

<!-- # Devices

Use the `devices` attribute to map devices from your host machine into your build container. These are [Docker devices](https://docs.docker.com/compose/compose-file/#devices) and therefore use the same declaration conventions:

```yaml
---
build:
  image: golang
  devices:
    - "/dev/ttyUSB0:/dev/ttyUSB0"
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen. -->

# Privileged

Use the `privileged` attribute to run your build in a privileged Docker container:

```yaml
---
build:
  image: golang
  privileged: true
```

For security reasons this option is only available to trusted repositories. Trusted repositories are enabled on a per-repository basis by a Drone administrator from your repository settings screen.

# Skip Commits

Instruct Drone to skip builds by by including any combination of `ci` and `skip` wrapped in square brackets in your commit message. Examples: `[skip CI]` `[ci skip]`

# Skip Branches

Instruct Drone to skip branches by including a branch white-list in your `.drone.yml`

```yaml
---
branches:
  - master
  - develop
```
