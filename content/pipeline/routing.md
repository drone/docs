---
date: 2000-01-01T00:00:00+00:00
title: Routing
author: bradrydzewski
weight: 1000
hide: true
draft: true
description: |
  Route pipelines to specific servers.
---

The `node` section can be used to route pipelines to specific runners, or groups of runners, that have matching labels. This can be useful when you need to route pipelines to runners with special configurations or hardware.

<div class="alert alert-warn">
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

node:
  keyA: valueA
  keyB: valueB
{{< / highlight >}}

# Kubernetes

The `node_selector` section should be used with Kubernetes pipelines to route workloads to specific nodes. Node selection can be used in conjunction with [taints and tolerations]({{< relref "/pipeline/kubernetes/configure/tolerations.md" >}}).

{{< highlight yaml "linenos=table,hl_lines=2 12-14" >}}
kind: pipeline
type: kubernetes
name: default

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test

node_selector:
  keyA: valueA
  keyB: valueB
{{< / highlight >}}