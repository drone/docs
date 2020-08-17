---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
description: |
  Overview of ssh pipelines.
---

An `ssh` pipeline executes shell commands on remote servers using the ssh protocol. This is useful for workloads that need to run directly on the host, or are poorly suited for execution inside containers.

Example pipeline configuration:

{{< highlight text "linenos=table" >}}
---
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

steps:
- name: greeting
  commands:
  - echo hello world

...
{{< / highlight >}}

The kind and type attributes define an ssh pipeline.

{{< highlight text "linenos=table" >}}
---
kind: pipeline
type: ssh
{{< / highlight >}}

The steps section defines a series of shell commands. These commands are executed on the remove server using the ssh protocol. If any command returns a non-zero exit code, the pipeline fails and exits.

{{< highlight text "linenos=table,linenostart=10" >}}
steps:
- name: greeting
  commands:
  - echo hello world
{{< / highlight >}}