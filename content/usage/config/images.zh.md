+++
date = "2017-04-15T14:39:04+02:00"
title = "Images 镜像"
url = "zh/images"

[menu.usage]
  weight = 1
  identifier = "images-zh"
  parent = "usage_config"
+++

<!--Drone uses Docker images for the build environment, for plugins and for service containers. The image field is exposed in the container blocks in the Yaml:-->

 在 Drone 中，构建环境、插件以及服务容器都是 Docker 镜像。在 Yaml 中使用 image 来指定使用的镜像。

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

<!--Drone supports any valid Docker image from any Docker registry:-->

Drone 支持来自任何 Docker registry 的任何有效的 Docker 镜像。

```text
image: golang
image: golang:1.7
image: library/golang:1.7
image: index.docker.io/library/golang
image: index.docker.io/library/golang:1.7
```
