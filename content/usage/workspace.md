+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Workspace"
weight = 2
menu = "usage"
toc = true
+++

# Overview

The project workspace refers to the location of your source code on disk, mounted into all containers that are part of your build process. All containers and plugins are started with the workspace and the working directory. This section of the documentation describes the default workspace configuration and customization options.

# Defaults

The default workspace configuration mounts the `/drone` volume into your containers. The workspace path, where your source repository is cloned, defaults to `src/{project}` relative to the base directory.

# Configuration

The default workspace can be customized using the `workspace` section of your `.drone.yml` file. In general we recommend using the default workspace, however, there are valid reasons for overriding default values. The official Go image, for example, expects the source code to exist in the `/go` directory.

Example custom base directory with default path:

```
workspace:
  base: /go

  # source code cloned to /go/src/{project}
```

Example default base directory with custom path:

```
workspace:
  path: src/github.com/octocat/hello-world

  # source code cloned to /drone/src/github.com/octocat/hello-world
```

Example custom base directory and custom path:

```
workspace:
  base: /go
  path: src/github.com/octocat/hello-world

  # source code cloned to /go/src/github.com/octocat/hello-world
```

# Examples

The below example alters the location of the project workspace to use the `/go` directory as the mount point, with the source cloned into a subdirectory. The official Go image is configured to compile code from the `/go` directory:

```
workspace:
  base: /go
  path: src/github.com/octocat/hello-world

script:
  build:
    image: golang:1.6
```

An alternative to the above workspace customization would be to override the default image configuration. Using the official Go image as an example, we can alter the GOPATH to use the default base directory.

```
script:
  build:
    image: golang:1.6
    environment:
      - GOPATH=/drone
```
