---
date: 2000-01-01T00:00:00+00:00
title: Parallelism
author: tphoney
weight: 90
toc: false
description: |
  Configure pipeline steps to execute in parallel.
---

Pipeline steps are executed sequentially by default. You can optionally describe your build steps as a [directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph). In the below example we fan-out to execute the first two steps in parallel, and then once complete, we fan-in to execute the final step:

{{<highlight yaml "linenos=table,hl_lines=23-25" >}}
kind: pipeline
type: aws
name: default

pool:
  use: ubuntu

steps:
- name: backend
  image: golang
  commands:
  - go build
  - go test

- name: frontend
  image: node
  commands:
  - npm install
  - npm test

- name: notify
  image: plugins/slack
  settings:
    webhook:
      from_secret: webhook
  depends_on:
  - frontend
  - backend
{{< / highlight >}}

The above example is quite simple, however, you can use this syntax to create very complex execution flows.

<div class="alert">
Note that when you define the dependency graph you must configure dependencies for all pipeline steps.
</div>

<div class="alert">
Note that you can use conditional steps in your dependency graph. The scheduler automatically corrects the dependency graph for skipped steps.
</div>
