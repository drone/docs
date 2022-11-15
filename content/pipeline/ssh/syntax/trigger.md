---
date: 2000-01-01T00:00:00+00:00
title: Triggers
author: bradrydzewski
weight: -1
toc: true
description: |
  Configure when pipeline executions are triggered.
---

When you push code to your repository, open a pull request, or create a tag, your source control management system automatically sends a webhook to Drone which in turn triggers pipeline execution. Use the triggers section to limit pipeline execution.

Example limits pipeline execution by branch:

```yaml {linenos=table, hl_lines=["16-19"]}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

steps:
- name: build
  commands:
  - go build
  - go test

trigger:
  branch:
  - master
```

You can use wildcard matching in your triggers. _Note that triggers use glob pattern matching, not regular expressions._

```yaml {linenos=table, linenostart=17}
trigger:
  ref:
  - refs/heads/master
  - refs/heads/**
  - refs/pull/*/head
```

You can also combine multiple triggers. _Note that all triggers must evaluate to true when combining multiple triggers._

```yaml {linenos=table, linenostart=17}
trigger:
  branch:
  - master
  event:
  - push
```

# By Branch

The branch trigger limits step execution based on the git branch. Please note that the target branch is evaluated for pull requests; and branch names are not available for tag events.

<div class="alert">
Note that you cannot use branch triggers with tags. A tag is not associated with the source branch from which it was created.
</div>

```yaml {linenos=table, linenostart=17}
trigger:
  branch:
  - master
  - feature/*
```

Example include syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  branch:
    include:
    - master
    - feature/*
```

Example exclude syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  branch:
    exclude:
    - master
    - feature/*
```

# By Event

The event trigger limits step execution based on the drone event type. This can be helpful when you want to limit steps based on push, pull request, tag and more.

<div class="alert">
Note that you cannot use branch triggers with tag events. A tag is not associated with the source branch from which it was created.
</div>

```yaml {linenos=table, linenostart=17}
trigger:
  event:
  - push
  - pull_request
  - tag
  - promote
  - rollback
```

Example include syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  event:
    include:
    - push
    - pull_request
```

Example exclude syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  event:
    exclude:
    - pull_request
```

# By Reference

The reference trigger limits step execution based on the git reference name. This can be helpful when you want to glob match branch or tag names.

```yaml {linenos=table, linenostart=17}
trigger:
  ref:
  - refs/heads/feature-*
  - refs/tags/*
```

Example include syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  ref:
    include:
    - refs/heads/feature-*
    - refs/pull/**
    - refs/tags/**
```

Example exclude syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  ref:
    exclude:
    - refs/heads/feature-*
    - refs/pull/**
    - refs/tags/**
```

# By Repository

The repository trigger limits step execution based on repository name. This can be useful when Drone is enabled for a repository and its forks, and you want to limit execution accordingly.

```yaml {linenos=table, linenostart=17}
trigger:
  repo:
  - octocat/hello-world
```

Example include syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  repo:
    include:
    - octocat/hello-world
    - spacebhost/hello-world
```

Example exclude syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  repo:
    exclude:
    - octocat/hello-world
    - spacebhost/hello-world
```

Example using wildcard matching:

```yaml {linenos=table, linenostart=17}
trigger:
  repo:
    include:
    - octocat/*
```

# By Instance

The instance trigger limits step execution based on the Drone instance hostname. This can be useful if you have multiple Drone instances configured for a single repository, sharing the same yaml file, and want to limit steps by instance.

```yaml {linenos=table, linenostart=17}
trigger:
  instance:
  - drone.instance1.com
  - drone.instance2.com
```

Example include syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  instance:
    include:
    - drone.instance1.com
    - drone.instance2.com
```

Example exclude syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  instance:
    exclude:
    - drone.instance1.com
    - drone.instance2.com
```

Example using wildcard matching:

```yaml {linenos=table, linenostart=17}
trigger:
  instance:
    include:
    - *.company.com
```

# By Status

The status trigger limits step execution based on the pipeline status. For example, you may want to configure Slack notification only on failure.

```yaml {linenos=table, linenostart=17}
trigger:
  status:
  - failure
```

Execute a step on failure:

```yaml {linenos=table, linenostart=17}
trigger:
  status:
  - failure
```

Execute a step on success or failure:

```yaml {linenos=table, linenostart=17}
trigger:
  status:
  - success
  - failure
```

The following configuration is redundant. The default behavior is for pipeline steps to only execute when the pipeline is in a passing state.

```yaml {linenos=table, linenostart=17}
trigger:
  status:
  - success
```

# By Target

The target trigger limits step execution based on the target deployment environment. This only applies to promotion and rollback events.

```yaml {linenos=table, linenostart=17}
trigger:
  target:
  - production
```

Example include syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  target:
    include:
    - staging
    - production
```

Example exclude syntax:

```yaml {linenos=table, linenostart=17}
trigger:
  target:
    exclude:
    - production
```