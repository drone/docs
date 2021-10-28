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

{{< highlight yaml "linenos=table,hl_lines=11-13" >}}
kind: pipeline
type: macstadium
name: default

steps:
- name: build
  commands:
  - go build
  - go test

trigger:
  branch:
  - master
{{< / highlight >}}

You can use wildcard matching in your triggers. _Note that triggers use glob pattern matching, not regular expressions._

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  ref:
  - refs/heads/master
  - refs/heads/**
  - refs/pull/*/head
{{< / highlight >}}

You can also combine multiple triggers. _Note that all triggers must evaluate to true when combining multiple triggers._

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  branch:
  - master
  event:
  - push
{{< / highlight >}}

# By Branch

The branch trigger limits step execution based on the git branch. Please note that the target branch is evaluated for pull requests; and branch names are not available for tag events.

<div class="alert">
Note that you cannot use branch triggers with tags. A tag is not associated with the source branch from which it was created.
</div>

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  branch:
  - master
  - feature/*
{{< / highlight >}}

Example include syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  branch:
    include:
    - master
    - feature/*
{{< / highlight >}}

Example exclude syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  branch:
    exclude:
    - master
    - feature/*
{{< / highlight >}}

# By Event

The event trigger limits step execution based on the drone event type. This can be helpful when you want to limit steps based on push, pull request, tag and more.

<div class="alert">
Note that you cannot use branch triggers with tag events. A tag is not associated with the source branch from which it was created.
</div>

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  event:
  - push
  - pull_request
  - tag
  - promote
  - rollback
{{< / highlight >}}

Example include syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  event:
    include:
    - push
    - pull_request
{{< / highlight >}}

Example exclude syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  event:
    exclude:
    - pull_request
{{< / highlight >}}

# By Reference

The reference trigger limits step execution based on the git reference name. This can be helpful when you want to glob match branch or tag names.

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  ref:
  - refs/heads/feature-*
  - refs/tags/*
{{< / highlight >}}

Example include syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  ref:
    include:
    - refs/heads/feature-*
    - refs/pull/**
    - refs/tags/**
{{< / highlight >}}

Example exclude syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  ref:
    exclude:
    - refs/heads/feature-*
    - refs/pull/**
    - refs/tags/**
{{< / highlight >}}

# By Repository

The repository trigger limits step execution based on repository name. This can be useful when Drone is enabled for a repository and its forks, and you want to limit execution accordingly.

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  repo:
  - octocat/hello-world
{{< / highlight >}}

Example include syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  repo:
    include:
    - octocat/hello-world
    - spacebhost/hello-world
{{< / highlight >}}

Example exclude syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  repo:
    exclude:
    - octocat/hello-world
    - spacebhost/hello-world
{{< / highlight >}}

Example using wildcard matching:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  repo:
    include:
    - octocat/*
{{< / highlight >}}

# By Instance

The instance trigger limits step execution based on the Drone instance hostname. This can be useful if you have multiple Drone instances configured for a single repository, sharing the same yaml file, and want to limit steps by instance.

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  instance:
  - drone.instance1.com
  - drone.instance2.com
{{< / highlight >}}

Example include syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  instance:
    include:
    - drone.instance1.com
    - drone.instance2.com
{{< / highlight >}}

Example exclude syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  instance:
    exclude:
    - drone.instance1.com
    - drone.instance2.com
{{< / highlight >}}

Example using wildcard matching:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  instance:
    include:
    - *.company.com
{{< / highlight >}}

# By Status

The status trigger limits step execution based on the pipeline status. For example, you may want to configure Slack notification only on failure.

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  status:
  - failure
{{< / highlight >}}

Execute a step on failure:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  status:
  - failure
{{< / highlight >}}

Execute a step on success or failure:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  status:
  - success
  - failure
{{< / highlight >}}

The following configuration is redundant. The default behavior is for pipeline steps to only execute when the pipeline is in a passing state.

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  status:
  - success
{{< / highlight >}}

# By Target

The target trigger limits step execution based on the target deployment environment. This only applies to promotion and rollback events.

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  target:
  - production
{{< / highlight >}}

Example include syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  target:
    include:
    - staging
    - production
{{< / highlight >}}

Example exclude syntax:

{{< highlight yaml "linenos=table,linenostart=11" >}}
trigger:
  target:
    exclude:
    - production
{{< / highlight >}}

# By Action

The action trigger limits execution based on the action associated with the event that triggered the pipeline.

<div class="alert">
Action support varies by event and SCM provider.
</div>

{{< highlight yaml "linenos=table,linenostart=12" >}}
trigger:
  action:
  - labeled
{{< / highlight >}}

Execute when a pull request is opened:

{{< highlight yaml "linenos=table,linenostart=12" >}}
trigger:
  event:
  - pull_request
  action:
  - opened
{{< / highlight >}}

Execute on all pull request actions except "synchronized":

{{< highlight yaml "linenos=table,linenostart=12" >}}
trigger:
  event:
  - pull_request
  action:
    exclude:
    - synchronized
{{< / highlight >}}
