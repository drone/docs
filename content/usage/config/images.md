+++
date = "2017-04-15T14:39:04+02:00"
title = "Images"
url = "images"

[menu.usage]
  weight = 1
  identifier = "images"
  parent = "usage_config"
+++

Drone uses Docker images for the build environment, for plugins and for service containers. The image field is exposed in the container blocks in the Yaml:

```diff
pipeline:
  build:
+   image: golang:1.6
    commands:
      - go build
      - go test

  publish:
+   image: plugins/docker
    repo: foo/bar

services:
  database:
+   image: mysql
```

Drone supports any valid Docker image from any Docker registry:

```text
image: golang
image: golang:1.7
image: library/golang:1.7
image: index.docker.io/library/golang
image: index.docker.io/library/golang:1.7
```
