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

description: |
  Overview of Kubernetes pipelines.
---

<div class="alert">
Please note Kubernetes pipelines are disabled on Drone Cloud. This feature is only available when self-hosting
</div>

A `kubernetes` pipeline executes pipeline steps as containers inside a Kubernetes Pod. Containers provide isolation allowing safe execution of concurrent pipelines on the same machine.

Example pipeline configuration:

{{< highlight text "linenos=table" >}}
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
{{< / highlight >}}

The kind and type attributes define a Kubernetes pipeline.

{{< highlight text "linenos=table" >}}
---
kind: pipeline
type: kubernetes
{{< / highlight >}}

The `steps` section defines a series of shell commands. These commands are executed inside the Docker container as the `Entrypoint`. If any command returns a non-zero exit code, the pipeline fails and exits.

{{< highlight text "linenos=table,linenostart=6" >}}
steps:
- name: greeting
  image: golang:1.12
  commands:
  - go build
  - go test
{{< / highlight >}}

To bind a non-default service account to a pipeline to allow for more fine grained access to resouces.

{{< highlight text "linenos=table" >}}
---
kind: pipeline
type: kubernetes
service_account_name: builder
{{< / highlight >}}
