---
date: 2000-01-01T00:00:00+00:00
title: Routing
author: bradrydzewski
weight: 120
toc: false
aliases:
- /user-guide/pipeline/nodes
description: |
  Route pipelines to specific servers.
---

The `nodes` section can be used to route pipelines to groups of runners with [matching labels]({{< relref "runner/docker/configuration/reference/drone-runner-labels.md" >}}). This can be useful when you need to route pipelines to runners with special configurations or hardware, for example a pool of runners with gpus or with high memory.

<div class="alert">
A pipeline is not routed to a runner unless it matches all runner labels. If the pipeline only defines and matches a subset of runner labels it will not be routed to the runner.
</div>

{{< highlight yaml "linenos=table,hl_lines=12-14" >}}
kind: pipeline
type: docker
name: default

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test

nodes:
  keyA: valueA
  keyB: valueB
{{< / highlight >}}
