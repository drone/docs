---
date: 2000-01-01T00:00:00+00:00
title: Parallelism
author: bradrydzewski
weight: 5
toc: false
description: |
  Configure pipeline steps to execute in parallel.
---

Pipeline steps are executed sequentially by default. You can optionally describe your build steps as a [directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph). In the below example we fan-out to execute the first two steps in parallel, and then once complete, we fan-in to execute the final step:

{{< highlight text "linenos=table,hl_lines=24-26" >}}
kind: pipeline
type: digitalocean
name: default

token:
  from_secret: token

steps:
- name: backend
  commands:
  - go build
  - go test

- name: frontend
  image: node
  commands:
  - npm install
  - npm test

- name: publish
  commands:
  - docker build -t hello-world .
  - docker push hello-world
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
