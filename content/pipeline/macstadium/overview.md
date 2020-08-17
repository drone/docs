---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1

description: |
  Overview of MacStadium pipelines.
---

<div class="alert">
Please note MacStadium pipelines are disabled on Drone Cloud. This feature is only available when self-hosting
</div>

A `macstadium` pipeline runs commands on MacStadium virtual machines. The vm is created when the pipeline starts and terminated upon completion. This is useful for workloads that need to run inside a dedicated osx vm.

Example pipeline configuration:

{{< highlight yaml "linenos=table" >}}
---
kind: pipeline
type: macstadium
name: default

steps:
- name: greeting
  commands:
  - echo hello world

...
{{< / highlight >}}

The kind and type attributes define a macstadium pipeline.

{{< highlight yaml "linenos=table" >}}
---
kind: pipeline
type: macstadium
{{< / highlight >}}

The `steps` section defines a series of shell commands. These commands are executed on the remote macstadium virtual machine. If any command returns a non-zero exit code, the pipeline fails and exits.

{{< highlight yaml "linenos=table,linenostart=9" >}}
steps:
- name: greeting
  commands:
  - echo hello world
{{< / highlight >}}