+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Deployments"
weight = 8
menu = "usage"
toc = true
+++

# Usage

Deployments

# Docker

Deployments are handled by plugins. We have multiple plugins for building and publishing Docker images to DockerHub, Google Container Engine, and more.

# Conditions

Use the `when` block to limit deployment step execution:

```
deploy:
  heroku:
    app: mycompany.com
    when:
      branch: master
```

Execute a deployment step if the branch is `master` or `develop`:

```
deploy:
  heroku:
    app: mycompany.com
    when:
      branch: [master, develop]
```

Execute a deployment step if the branch is _not_ `master`:

```
deploy:
  heroku:
    app: mycompany.com
    when:
      branch: "!master"
```

Execute a deployment step if the branch is starts with `prefix/*`:

```
deploy:
  heroku:
    app: mycompany.com
    when:
      branch: prefix/*
```

Execute a deployment step when the commit is a `tag`:

```
deploy:
  heroku:
    app: mycompany.com
    when:
      event: tag
```

Execute a deployment step for a single matrix permutation:

```
deploy:
  heroku:
    app: mycompany.com
    when:
      matrix:
        GO_VERSION: 1.5
        REDIS_VERSION: 2.8
```

Define the same deployment step multiple times, using different configuration based on branch:

```
deploy:
  heroku:
    app: test.mycompany.com
    when:
      branch: develop

  heroku:
    app: mycompany.com
    when:
      branch: master
```
