+++
date = "2017-04-15T14:39:04+02:00"
title = "Matrix Builds 矩阵构建"
url = "zh/matrix-builds"

[menu.usage]
  weight = 10
  identifier = "matrix-builds-zh"
  parent = "usage_concepts"
+++


<!--Drone has integrated support for matrix builds. Drone executes a separate build task for each combination in the matrix, allowing you to build and test a single commit against multiple configurations.-->

Drone 支持整合矩阵构建。Drone 为配置文件矩阵中的每个组合各执行一个独立的构建任务。矩阵构建允许您使用多个配置来构建和测试一个提交（commit）。

<!--Example matrix definition:-->

矩阵定义示例：

```yaml
matrix:
  GO_VERSION:
    - 1.4
    - 1.3
  REDIS_VERSION:
    - 2.6
    - 2.8
    - 3.0
```

<!--Example matrix definition containing only specific combinations:-->

只包含特定组合的矩阵示例：

```yaml
matrix:
  include:
    - GO_VERSION: 1.4
      REDIS_VERSION: 2.8
    - GO_VERSION: 1.5
      REDIS_VERSION: 2.8
    - GO_VERSION: 1.6
      REDIS_VERSION: 3.0
```

# 变量插值

<!--Matrix variables are interpolated in the yaml using the `${VARIABLE}` syntax, before the yaml is parsed. This is an example yaml file before interpolating matrix parameters:-->

矩阵变量可以使用 `${VARIABLE}` 的语法来进行插值。

```yaml
pipeline:
  build:
    image: golang:${GO_VERSION}
    commands:
      - go get
      - go build
      - go test

services:
  database:
    image: ${DATABASE}

matrix:
  GO_VERSION:
    - 1.4
    - 1.3
  DATABASE:
    - mysql:5.5
    - mysql:6.5
    - mariadb:10.1
```

<!--Example Yaml file after injecting the matrix parameters:-->

插入矩阵变量后的 Yaml 文件示例：

```diff
pipeline:
  build:
-   image: golang:${GO_VERSION}
+   image: golang:1.4
    commands:
      - go get
      - go build
      - go test
+   environment:
+     - GO_VERSION=1.4
+     - DATABASE=mysql:5.5

services:
  database:
-   image: ${DATABASE}
+   image: mysql:5.5
```

# 示例

<!--Example matrix build based on Docker image tag:-->

基于 Docker 镜像标签（Docker image tag）的矩阵构建

```yaml
pipeline:
  build:
    image: golang:${TAG}
    commands:
      - go build
      - go test

matrix:
  TAG:
    - 1.7
    - 1.8
    - latest
```

<!--Example matrix build based on Docker image:-->

基于 Docker 镜像的矩阵构建

```yaml
pipeline:
  build:
    image: ${IMAGE}
    commands:
      - go build
      - go test

matrix:
  IMAGE:
    - golang:1.7
    - golang:1.8
    - golang:latest
```
