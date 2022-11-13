---
date: 2000-01-01T00:00:00+00:00
title: Conditions
author: bradrydzewski
weight: 7
toc: true
description: |
  Configure pipeline steps to execute conditionally.
---

Conditions can be used to limit pipeline step execution at runtime. For example, you may want to limit step execution by branch:

```yaml {linenos=table, hl_lines=["10-13"]}
kind: pipeline
type: exec
name: default

steps:
- name: build
  commands:
  - go build
  - go test
  when:
    branch:
    - master
    - feature/*
```

You can use wildcard matching in your conditions. _Note that conditions use [glob](https://golang.org/pkg/path/#Match) pattern matching, not regular expressions._

```yaml {linenos=table, linenostart=10}
when:
  ref:
  - refs/heads/master
  - refs/heads/**
  - refs/pull/*/head
```

You can also combine multiple conditions. _Note that all conditions must evaluate to true when combining multiple conditions._

```yaml {linenos=table, linenostart=10}
when:
  branch:
  - master
  event:
  - push
```

# By Branch

The branch condition limits step execution based on the git branch. Please note that the target branch is evaluated for pull requests; and branch names are not available for tag events.

<div class="alert">
Note that you cannot use branch conditions with tags. A tag is not associated with the source branch from which it was created.
</div>

```yaml {linenos=table, linenostart=10}
when:
  branch:
  - master
  - feature/*
```

Example include syntax:

```yaml {linenos=table, linenostart=10}
when:
  branch:
    include:
    - master
    - feature/*
```

Example exclude syntax:

```yaml {linenos=table, linenostart=10}
when:
  branch:
    exclude:
    - master
    - feature/*
```

# By Event

The event condition limits step execution based on the drone event type. This can be helpful when you want to limit steps based on push, pull request, tag and more.

<div class="alert">
Note that you cannot use branch conditions with tag events. A tag is not associated with the source branch from which it was created.
</div>

```yaml {linenos=table, linenostart=10}
when:
  event:
  - push
  - pull_request
  - tag
  - promote
  - rollback
```

Example include syntax:

```yaml {linenos=table, linenostart=10}
when:
  event:
    include:
    - push
    - pull_request
```

Example exclude syntax:

```yaml {linenos=table, linenostart=10}
when:
  event:
    exclude:
    - pull_request
```

# By Reference

The reference condition limits step execution based on the git reference name. This can be helpful when you want to glob match branch or tag names.


```yaml {linenos=table, linenostart=10}
when:
  ref:
  - refs/heads/feature-*
  - refs/tags/*
```

Example include syntax:

```yaml {linenos=table, linenostart=10}
when:
  ref:
    include:
    - refs/heads/feature-*
    - refs/pull/**
    - refs/tags/**
```

Example exclude syntax:

```yaml {linenos=table, linenostart=10}
when:
  ref:
    exclude:
    - refs/heads/feature-*
    - refs/pull/**
    - refs/tags/**
```

# By Repository

The repository condition limits step execution based on repository name. This can be useful when Drone is enabled for a repository and its forks, and you want to limit execution accordingly.


```yaml {linenos=table, linenostart=10}
when:
  repo:
  - octocat/hello-world
```

Example include syntax:

```yaml {linenos=table, linenostart=10}
when:
  repo:
    include:
    - octocat/hello-world
    - spacebhost/hello-world
```

Example exclude syntax:

```yaml {linenos=table, linenostart=10}
when:
  repo:
    exclude:
    - octocat/hello-world
    - spacebhost/hello-world
```

Example using wildcard matching:

```yaml {linenos=table, linenostart=10}
when:
  repo:
    include:
    - octocat/*
```

# By Instance

The instance condition limits step execution based on the Drone instance hostname. This can be useful if you have multiple Drone instances configured for a single repository, sharing the same yaml file, and want to limit steps by instance.

```yaml {linenos=table, linenostart=10}
when:
  instance:
  - drone.instance1.com
  - drone.instance2.com
```

Example include syntax:

```yaml {linenos=table, linenostart=10}
when:
  instance:
    include:
    - drone.instance1.com
    - drone.instance2.com
```

Example exclude syntax:

```yaml {linenos=table, linenostart=10}
when:
  instance:
    exclude:
    - drone.instance1.com
    - drone.instance2.com
```

Example using wildcard matching:

```yaml {linenos=table, linenostart=10}
when:
  instance:
    include:
    - *.company.com
```

# By Status

The status condition limits step execution based on the pipeline status. For example, you may want to configure Slack notification only on failure.

```yaml {linenos=table, linenostart=10}
when:
  status:
  - failure
```

Execute a step on failure:

```yaml {linenos=table, linenostart=10}
when:
  status:
  - failure
```

Execute a step on success or failure:

```yaml {linenos=table, linenostart=10}
when:
  status:
  - success
  - failure
```

The following configuration is redundant. The default behavior is for pipeline steps to only execute when the pipeline is in a passing state.

```yaml {linenos=table, linenostart=10}
when:
  status:
  - success
```

# By Target

The target condition limits step execution based on the target deployment environment. This only applies to promotion and rollback events.

```yaml {linenos=table, linenostart=10}
when:
  target:
  - production
```

Example include syntax:

```yaml {linenos=table, linenostart=10}
when:
  target:
    include:
    - staging
    - production
```

Example exclude syntax:

```yaml {linenos=table, linenostart=10}
when:
  target:
    exclude:
    - production
```