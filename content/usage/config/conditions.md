+++
date = "2017-04-15T14:39:04+02:00"
title = "Conditions"
url = "conditional-steps"

[menu.usage]
  weight = 2
  identifier = "conditions"
  parent = "usage_config"
+++

Drone support defining condition pipeline steps in the `when` block. If all conditions in the `when` block evaluate to true the step is executed, otherwise it is skipped.

Example conditional execution by branch:

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
+   when:
+     branch: master
```

# Branches

Execute a step if the branch is `master` or `develop`:

```diff
when:
  branch: [master, develop]
```

Execute a step if the branch is starts with `prefix/*`:

```diff
when:
  branch: prefix/*
```

Execute a step using custom include and exclude logic:

```diff
when:
  branch:
    include: [ master, release/* ]
    exclude: [ release/1.0.0, release/1.1.* ]
```   

# Events

Execute a step if the build event is a `tag`:

```diff
when:
  event: tag
```

Execute a step for all non-pull request events:

```diff
when:
  event: [push, tag, deployment]
```

Execute a step for all build events:

```diff
when:
  event: [push, pull_request, tag, deployment]
```

# Status

Execute a step when the build status changes:

```diff
when:
  status: changed
```

Execute a step when the build is passing or failing:

```diff
when:
  status:  [ failure, success ]
```

# Platform

Execute a step for a specific platform:

```diff
when:
  platform: linux/amd64
```

Execute a step for a specific platform using wildcards:

```diff
when:
  platform:  [ linux/*, windows/amd64 ]
```

# Environment

Execute a step for deployment events matching the target deployment environment:

```diff
when:
  environment: production
  event: deployment
```

# Matrix

Execute a step for a single matrix permutation:

```diff
when:
  matrix:
    GO_VERSION: 1.5
    REDIS_VERSION: 2.8
```
