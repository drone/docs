---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1

description: |
  Overview of Digital Ocean pipelines.
---

<div class="alert">
Please note Digital Ocean pipelines are disabled on Drone Cloud. This feature is only available when self-hosting
</div>

A `digitalocean` pipeline runs shell commands on digitalocean droplets. The droplet is created when the pipeline starts and terminated upon completion. This is useful for workloads that need to run inside a dedicated virtual machine.

Example pipeline configuration:

```yaml {linenos=table}
---
kind: pipeline
type: digitalocean
name: default

token:
  from_secret: token

steps:
- name: greeting
  commands:
  - echo hello world

...
```

The kind and type attributes define a digitalocean pipeline.

```yaml {linenos=table}
---
kind: pipeline
type: digitalocean
```

The `steps` section defines a series of shell commands. These commands are executed on the remote digital ocean droplet. If any command returns a non-zero exit code, the pipeline fails and exits.

```yaml {linenos=table, linenostart=9}
steps:
- name: greeting
  commands:
  - echo hello world
```