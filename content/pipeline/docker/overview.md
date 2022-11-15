---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1

related:
  items:
  - name: Configure Pipeline Steps
    path: syntax/steps.md
  - name: Configure Pipeline Services
    path: syntax/services.md
  - name: Integrations using Plugins
    path: syntax/plugins.md

aliases:
- /yaml-reference/
- /user-guide/pipeline

description: |
  Overview of Docker pipelines.
---

Docker pipelines execute pipeline commands inside ephemeral Docker containers. Docker containers provide isolation, allowing safe execution of concurrent pipelines on the same machine.

Example pipeline configuration:

```yaml {linenos=table}
---
kind: pipeline
type: docker
name: default

steps:
- name: greeting
  image: golang:1.12
  commands:
  - go build
  - go test

...
```

The kind and type attributes define a Docker pipeline.

```yaml {linenos=table}
---
kind: pipeline
type: docker
```

The `steps` section defines a series of shell commands. These commands are executed inside the Docker container as the `Entrypoint`. If any command returns a non-zero exit code, the pipeline fails and exits.

```yaml {linenos=table, linenostart=6}
steps:
- name: greeting
  image: golang:1.12
  commands:
  - go build
  - go test
```