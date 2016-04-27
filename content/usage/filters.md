+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Filters"
weight = 28
menu = "usage"
toc = true
+++

# Overview

You can use the `when` block to limit the execution of any step in the build process. This includes command execution steps and plugin steps. The below example demonstrates filter logic based on branch:


```yaml
script:
  ...

  heroku:
    app: mycompany.com
    when:
      branch: master
```

# Branches

Execute a step if the branch is `master` or `develop`:

```yaml
when:
  branch: [master, develop]
```

Execute a step if the branch is starts with `prefix/*`:

```yaml
when:
  branch: prefix/*
```

# Events

Execute a step if the build event is a `tag`:

```yaml
when:
  event: tag
```

Execute a step if the build event matches one of the values:

```yaml
when:
  event: [push, pull_request, tag, deployment]
```

Execute a step if the build event matches one of the values:

```yaml
when:
  event: [push, pull_request, tag, deployment]
```

# Status

Execute a step when the build status changes:

```yaml
when:
  status: changed
```

Execute a step when the build is passing or failing:

```yaml
when:
  status:  [ failure, success ]
```

# Matrix

Execute a step for specific matrix steps:

```yaml
when:
  environment: production
  event: deployment
```

# Platform

Execute a step for a specific platform:

```yaml
when:
  platform: linux/amd64
```

Execute a step for a specific platform using wildcards:

```yaml
when:
  platform:  [ linux/*, windows/amd64 ]
```

# Environment

Execute a step for deployment events matching the target environment:

```yaml
when:
  environment: production
  event: deployment
```
