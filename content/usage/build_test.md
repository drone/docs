+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Build & Test"
weight = 7
menu = "usage"
toc = true
+++

# Overview

Drone uses the `build` section of the `.drone.yml` to define your Docker build environment and your build and test instructions. The following is an example build definition:

```
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

```
image: golang
image: golang:1.5
image: library/golang:1.5
image: index.docker.io/library/golang:1.5
```

Provide your registry credentials if your build image is private:

```
build:
  image: foo/bar
  auth_config:
    username: octocat
    password: password
```

Use the `pull` attribute to instruct Drone to always pull the latest Docker image. This helps ensure you are always testing your code against the latest image:

```
build:
  image: golang
  pull: true
```

# Workspace

The build workspace is located in the `/drone/src/<repository>` directory inside the `/drone` volume. You cannot change the volume location, however, you can change the subdirectory.

The below example overrides the default workspace to `/drone/src/github.com/octocat/hello-world`:
```
clone:
  path: github.com/octocat/hello-world
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

Please be aware that chaining commands together with `&&` may result in the build not exiting on failure. This is not a bug with Drone. This is the default bash `-e` behavior. See [this stackoverflow](http://stackoverflow.com/questions/25794905/why-does-set-e-true-false-true-not-exit#25795239) for more details.

# Dependencies

Cloning private submodules or private dependencies (using `npm` or `go get`) requires remote authentication. We recommend using `git+https` for submodules or dependencies, allowing Drone to authenticate to your remote repository using a `.netrc` file injected by default into your build environment.

Instead of using the `git+ssh` url in your `package.json` file:

```
{
    "library": "git+ssh://git@github.com:octocat/library.git"
}
```

Use the `git+https` url in your `package.json` file:

```
{
    "library": "https://github.com/octocat/library.git"
}
```

# Volumes

Use the `volumes` attribute to mount folders on your host machine into your build container. These are [Docker volumes](https://docs.docker.com/engine/userguide/dockervolumes/) and therefore use the same `<host>:<container>` declaration conventions:

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

# Triggers

Builds are triggered primarily by `push` and `pull_request` hooks from your remote system (ie GitHub). Sometimes changes to an upstream project have a cascading impact on downstream projects. You can trigger cascading builds using the [trigger](#) plugin.

# Scheduling

Drone does not have any builtin scheduling capabilities. For scheduled builds (ie nightly builds) we recommend using the Drone command line tools with cron to trigger builds at scheduled intervals.

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
