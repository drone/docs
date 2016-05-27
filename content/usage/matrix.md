+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Matrix"
weight = 29
toc = true
break = true

[menu.main]
	parent="usage"
+++

# Overview

Drone has integrated support for matrix builds. Drone executes a separate build task for each combination in the matrix, allowing you to build and test a single commit against multiple configurations.

Example matrix definition:

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

Example matrix definition containing only specific combinations:

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

# Interpolation

Matrix variables are interpolated in the yaml using the `${VARIABLE}` syntax, before the yaml is parsed. This is an example yaml file before interpolating matrix parameters:

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

Example Yaml file after injecting the matrix parameters:

```yaml
pipeline:
  build:
    image: golang:1.4
    environment:
      - GO_VERSION=1.4
      - DATABASE=mysql:5.5
    commands:
      - go get
      - go build
      - go test

services:
  database:
    image: mysql:5.5
```

# Conditions

Matrix builds execute the same Yaml multiple times but with different parameters. If you are using notification or deployment plugins you probably want to prevent multiple executions. Add the following condition to restrict a step to a single permutation:

```yaml
heroku:
  app: octokit
  force: false
  when:
    matrix:
      GO_VERSION: 1.4
      REDIS_VERSION: 3.0
```
