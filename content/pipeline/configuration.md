---
date: 2000-01-01T00:00:00+00:00
title: Configuration
author: bradrydzewski
weight: 3
toc: true
aliases:
- /configure/pipeline/multiple
- /configure/pipeline/multiple/
- /config/pipeline/multi-machine/
- /config/pipeline/multi-platform/
- /user-guide/pipeline/multi-machine
- /user-guide/pipeline/multi-platform
description: |
  Pipeline configuration overview.
---

This document provides a high-level overview of the .drone.yml configuration file. The configuration file is placed in the root of your repository and defines one or more continuous integration or continuous delivery [Pipelines]({{< relref "overview.md" >}}).

Example pipeline configuration:

```yaml {linenos=table}
kind: pipeline
type: docker
name: default

steps:
- name: test
  image: golang
  commands:
  - go build
  - go test
```

# Objects

The configuration file defines one or many objects in the same file, where each object is a separate yaml document. Pipelines are one kind of object. You can also define [Signature]({{< relref "signature/_index.md" >}}) and [Secret]({{< relref "secret/encrypted.md" >}}) objects.

* Example Pipeline configuration:
  ```
  kind: pipeline
  type: docker
  name: default

  steps:
  - name: test
    image: golang
    commands:
    - go build
    - go test
  ```

* Example Pipeline and Signature configuration:
    ```
    ---
    kind: pipeline
    type: docker
    name: default

    steps:
    - name: test
        image: golang
        commands:
        - go build
        - go test

    ---
    kind: signature
    hmac: F10E2821BBBEA527EA02200352313BC059445190
    ```

Each object must include the _kind_ attribute. The kind attribute identifies the kind of object being declared, and helps to distinguish one object from another.

# Pipelines

The _Pipeline_ object defines a Continuous Integration and Continuous Delivery pipeline. The _type_ attribute defines the runtime that should be used when executing the pipeline. Drone offers support for a variety of runtimes.

* Example Docker pipeline, which executes each pipeline steps inside isolated Docker containers.
  ```
  kind: pipeline
  type: docker
  ```

* Example Kubernetes pipeline, which executes pipeline steps as containers inside of Kubernetes pods.
  ```
  kind: pipeline
  type: kubernetes
  ```

* Example Exec pipeline, which executes pipeline steps directly on the host machine, with zero isolation.
  ```
  kind: pipeline
  type: exec
  ```

## Multiple Pipelines

When you define multiple _Pipeline_ objects in the same configuration file, pipelines are spread across runners and executed in parallel. The below example configures two pipelines to execute in parallel. The overall build status is determined by the successful completion of both pipelines.

<div class="alert">
Please note that Pipelines do not share state. It is not possible for two pipelines to access the same underlying file system or generated files.
</div>

<div class="alert">
When running multiple pipelines if any of the sibling pipelines fails further build pipelines will be stopped. 
This may present different behaviour, depending on the scheduling of sibling pipelines. This will give non-deterministic behaviour. 
</div>

{{< highlight yaml "linenos=table,hl_lines=35-57" >}}
kind: pipeline
type: docker
name: backend

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test

---
kind: pipeline
type: docker
name: frontend

steps:
- name: build
  image: node
  commands:
  - npm install
  - npm test
{{< / highlight >}}

## Multiple Platforms

One common use case for multiple pipelines is to define and execute pipelines for different platforms. For example, execute one pipeline on x86 and another pipeline on arm.

{{< highlight yaml "linenos=table,hl_lines=5-6 20-21" >}}
kind: pipeline
type: docker
name: amd64

platform:
  arch: amd64

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test

---
kind: pipeline
type: docker
name: arm

platform:
  arch: arm64

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}


## Graph Execution

Drone also supports describing your pipelines as a [directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph). In the below example we fan-out to execute the first two pipelines in parallel, and then once complete, we fan-in to execute the final pipeline.

{{< highlight yaml "linenos=table,hl_lines=35-57" >}}
kind: pipeline
type: docker
name: backend

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test

---
kind: pipeline
type: docker
name: frontend

steps:
- name: build
  image: node
  commands:
  - npm install
  - npm test

---
kind: pipeline
name: after

steps:
- name: notify
  image: plugins/slack
  settings:
    room: general
    webhook: https://...

depends_on:
- backend
- frontend
{{< / highlight >}}

# Scripting

Drone supports custom language extensions for when you need a more powerful alternative to yaml. See the [Starlark]({{< relref "scripting/starlark.md" >}}) and [Jsonnet]({{< relref "scripting/jsonnet.md" >}}) documentation for more details. You can also create your own language extension.
