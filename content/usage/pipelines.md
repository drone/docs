+++
date = "2015-12-05T16:00:21-08:00"
draft = true
title = "Pipelines"
weight = 29
menu = "usage"
toc = true
break = true
+++

# Overview

Drone supports simple yet robust build pipelines. It is important to note there are existing open source tools specializing in highly complex and flexible pipelines. If you find these existing tools complex or over-engineered we highly recommend Drone pipelines.



The actual deployment process and target environments are flexible


# Build and Test

The first part of the deployment pipeline is building and testing your code. After our tests are successful we build and publish a Docker image, and deploy that image to our development environment in Docker Cloud (Tutum).

```
script:
  test:
    image: golang
    commands:
      - go get
      - go build
      - go test
    when:
      event: [pull_request, push]

  docker:
    image: docker
    repo: octocat/hello-world
    tag: ${DRONE_COMMIT}
    when:
      event: push
      branch: master

  deploy:
    service: dev.app.com
    image: octocat/hello-world:${DRONE_COMMIT}
    when:
      event: push
      branch: master
```



```
script:
  ...


    image: golang
    commands:
      - go get
      - go build
      - go test
```



```
drone deploy {repository} {build} {environment}
```


# Promote To Stage

```
drone deploy octocat/hello-world 12 staging
```

```
script:
  ...

  stage:
    service: staging.app.com
    image: octocat/hello-world:${DRONE_COMMIT}
    when:
      event: deployment
      environment: staging
```

# Promote To Prod

You can

```
script:
  ...

  prod:
    service: app.com
    image: octocat/hello-world:${DRONE_COMMIT}
    when:
      event: deployment
      environment: production
```

```
drone deploy octocat/hello-world 12 production
```
