---
date: 2000-01-01T00:00:00+00:00
title: Steps
author: bradrydzewski
weight: 4
toc: true
description: |
  Configure pipeline steps.
---

Pipeline steps are defined as a series of shell commands. The commands are executed inside the root directory of your git repository. The root of your git repository, also called the workspace, is shared by all steps in your pipeline.

Example configuration:

{{< highlight yaml "linenos=table" >}}
kind: pipeline
type: digitalocean
name: default

steps:
- name: backend
  commands:
  - go build
  - go test

- name: frontend
  commands:
  - npm install
  - npm test
{{< / highlight >}}

# Commands

The commands are executed inside the root directory of your git repository. The root of your git repository, also called the workspace, is shared by all steps in your pipeline. This allows file artifacts to persist between steps.

{{< highlight yaml "linenos=table,linenostart=5" >}}
steps:
- name: backend
  commands:
  - go build
  - go test
{{< / highlight >}}

The above commands are converted to a simple shell script. The commands in the above example are roughly converted to the below script:

{{< highlight text "linenos=table" >}}
#!/bin/sh
set -e
set -x

go build
go test
{{< / highlight >}}

The exit code is used to determine whether the step is passing or failing. If a command returns a non-zero exit code, the step is marked as failing. The overall pipeline status is also marked as failing, and remaining pipeline steps are skipped (_unless explicitly configured to run on failure_).

# Environment

The environment section provides the ability to define environment variables scoped to individual pipeline steps.

{{< highlight yaml "linenos=table,hl_lines=3-5,linenostart=5" >}}
steps:
- name: backend
  environment:
    GOOS: linux
    GOARCH: amd64
  commands:
  - go build
  - go test
{{< / highlight >}}

<!--
See the Environment article for additional details:
TODO
-->

# Conditions

The when section provides the ability to conditionally limit the execution of steps at runtime. The below example limits step execution by branch, however, you can limit execution by event, reference, status and more.

{{< highlight yaml "linenos=table,hl_lines=6-8,linenostart=5" >}}
steps:
- name: backend
  commands:
  - go build
  - go test
  when:
    branch:
    - master
{{< / highlight >}}

Use the status condition to override the default runtime behavior and execute steps even when the pipeline status is failure:

{{< highlight yaml "linenos=table,hl_lines=5-8,linenostart=5" >}}
steps:
- name: cleanup
  commands:
  - docker system prune -f
  when:
    status:
    - failure
    - success
{{< / highlight >}}

See the Conditions article for additional details:

{{< link "/pipeline/macstadium/syntax/conditions.md" >}}

# Failure

The failure attribute lets you customize how the system handles failure of an individual step. This can be useful if you want to allow a step to fail without failing the overall pipeline.

{{< highlight yaml "linenos=table,hl_lines=3,linenostart=5" >}}
steps:
- name: backend
  failure: ignore
  commands:
  - go build
  - go test
{{< / highlight >}}
