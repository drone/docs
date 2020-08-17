---
date: 2000-01-01T00:00:00+00:00
title: Routing
author: bradrydzewski
weight: 9
toc: false
draft: true
---

The `nodes` section can be used to route pipelines to specific runners, or groups of runners, that have matching labels. This can be useful when you need to route pipelines to runners with special configurations or hardware.

{{< alert "warn" >}}
A pipeline is not routed to a runner unless it matches all runner labels. If the pipeline only defines and matches a subset of runner labels it will not be routed to the runner.
{{< / alert >}}

{{< highlight text "linenos=table,hl_lines=10-13" >}}
kind: pipeline
type: digitalocean
name: default

steps:
- name: build
  commands:
  - go build
  - go test

node:
  keyA: valueA
  keyB: valueB
{{< / highlight >}}
