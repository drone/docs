+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Matrix"
weight = 29
menu = "usage"
toc = true
break = true
+++

# Overview

Drone supports parallel matrix execution. Drone executes a separate build for each permutation in the matrix, allowing you to build and test a single commit against many configurations.

Example matrix configuration for testing multiple Go and Redis versions resulting:

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

Matrix variables are injected into the yaml using the `${VARIABLE}` syntax. Matrix variables are also available to your build container as environment variables.

Example Yaml file before interpolating matrix parameters:

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

# Includes

In some cases you may need more control over matrix permutations. You can use the `include` attribute to enumerate the available matrix permutations instead of having them calculated automatically.

```
matrix:
  include:
    - GO_VERSION: 1.4
      REDIS_VERSION: 2.8
    - GO_VERSION: 1.5
      REDIS_VERSION: 2.8
    - GO_VERSION: 1.6
      REDIS_VERSION: 3.0
```

# Conditions

Matrix builds execute the same Yaml multiple times, but with different parameters. If you are using notification or deployment plugins you probabaly want to prevent multiple executions. To restrict a step to a single permutation you can add the following condition:

```yaml
heroku:
  app: octokit
  force: false
  when:
    matrix:
      GO_VERSION: 1.4
      REDIS_VERSION: 3.0
```
