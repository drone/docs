---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
description: |
  Overview of exec pipelines.
---

<div class="alert">
Please note exec pipelines are disabled on Drone Cloud. This feature is only available when self-hosting
</div>

An `exec` pipeline executes shell commands directly on the host machine without isolation. This is useful for workloads that need to run on the host, or are poorly suited for execution inside containers.

Example pipeline configuration:

```yaml {linenos=table}
---
kind: pipeline
type: exec
name: default

platform:
  os: linux
  arch: amd64

steps:
- name: greeting
  commands:
  - echo hello world

...
```

The kind and type attributes define an exec pipeline.

```yaml {linenos=table}
---
kind: pipeline
type: exec
```

The platform section configures the target operating system and architecture, and ensures the pipeline is routed to the appropriate instance:

```yaml {linenos=table, linenostart=6}
platform:
  os: linux
  arch: amd64
```

The steps section defines a series of shell commands. These commands are executed using the default shell on posix, and powershell on windows. If any command returns a non-zero exit code, the pipeline fails and exits.

```yaml {linenos=table, linenostart=10}
steps:
- name: greeting
  commands:
  - echo hello world
```
