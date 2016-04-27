+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Matrix"
weight = 11
menu = "usage"
toc = true
+++

# Overview

Drone uses the `matrix` section of the `.drone.yml` to define the build matrix. Drone executes a build for each permutation in the matrix, allowing you to build and test a single commit against many configurations.

Example matrix definition for testing multiple Go and Redis versions. This matrix results in a total of 6 different build permutations:

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

# Interpolation

Matrix variables are injected into the `.drone.yml` file using the `${PARAM}` syntax, performing a simple find / replace. Matrix variables are also injected into your build container as environment variables.

Example Yaml file before injecting the matrix parameters:

```yaml
script:
  build:
    image: golang:${GO_VERSION}
    commands:
      - go get
      - go build
      - go test

services:
  redis:
    image: redis:${REDIS_VERSION}

matrix:
  GO_VERSION:
    - 1.4
    - 1.3
  REDIS_VERSION:
    - 2.6
    - 2.8
    - 3.0
```

Example Yaml file after injecting the matrix parameters:

```yaml
script:
  build:
    image: golang:1.4
    environment:
      - GO_VERSION=1.4
      - REDIS_VERSION=3.0
    commands:
      - go get
      - go build
      - go test

services:
  redis:
    image: redis:3.0
```

# Plugins

Matrix builds execute the same Yaml multiple times, but with different parameters. This means plugins are executed multiple times as well which may not be desired. To restrict a step to a single permutation you can add the following condition:

```yaml
heroku:
  app: octokit
  force: false
  when:
    matrix:
      GO_VERSION: 1.4
      REDIS_VERSION: 3.0
```
