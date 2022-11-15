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
  - name: DRONE_SERVICE_ACCOUNT_DEFAULT
    path: /runner/kubernetes/configuration/reference/drone-service-account-default.md

description: |
  Overview of Kubernetes pipelines.
---

<div class="alert">
Please note Kubernetes pipelines are disabled on Drone Cloud. This feature is only available when self-hosting
</div>

A `kubernetes` pipeline executes pipeline steps as containers inside a Kubernetes Pod. Containers provide isolation allowing safe execution of concurrent pipelines on the same machine.

Example pipeline configuration:

```yaml {linenos=table}
---
kind: pipeline
type: kubernetes
name: default

steps:
- name: greeting
  image: golang:1.12
  commands:
  - go build
  - go test

...
```

The kind and type attributes define a Kubernetes pipeline.

```yaml {linenos=table}
---
kind: pipeline
type: kubernetes
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

To bind a non-default service account to a pipeline and allow for more fine-grained access to resouces, do the following:

```yaml {linenos=table}
---
kind: pipeline
type: kubernetes
service_account_name: builder
```
