---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: tphoney
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
  Overview of AWS pipelines.
---

<div class="alert">
The AWS runner is in the Release Candidate phase.
</div>

<div class="alert">
Please note AWS pipelines are disabled on Drone Cloud. This feature is only available when self-hosting
</div>

An `aws` pipeline is executed on a dedicated, remote EC2 instance. The EC2 instance is created when the pipeline starts and terminated upon completion. This is useful for workloads that need to run inside a dedicated virtual machine.

Example pipeline configuration:

{{<highlight yaml "linenos=table" >}}
---
kind: pipeline
type: vm
name: default

pool:
  use: ubuntu-20.04

steps:
- name: test
  image: golang:1.12
  commands:
  - go test

- name: build
  commands:
  - docker build .
...
{{< / highlight >}}

## Windows support

AWS pipelines support windows EC2 instances as a first class citizen. Syntactically they do not differ to non Windows pipelines, but the underlying infrastructure is different.

{{<highlight yaml "linenos=table" >}}
---
kind: pipeline
type: vm
name: default

pool:
  use: windows-2019

steps:
- name: check install
  commands:
  - type C:\ProgramData\Amazon\EC2-Windows\Launch\Log\UserdataExecution.log
...
{{< / highlight >}}

## Kind and Type

The kind and type attributes define an `aws` pipeline.

{{<highlight yaml "linenos=table" >}}
---
kind: pipeline
type: vm
{{< / highlight >}}

## Pool

The pool attribute specifies the type of machine the pipeline should run on. _Please contact your system administrator for a list of machine types that are supported by your installation._

{{<highlight yaml "linenos=table,linenostart=5" >}}
pool:
  use: ubuntu-20.04
{{< / highlight >}}

For more information about configuring pools of machines please consult the [Pool file]({{< relref "../../runner/aws/configuration/pool.md" >}}) documentation.

## Steps

The `steps` section defines a series of shell commands. If any command returns a non-zero exit code, the pipeline fails and exits. Pipeline steps can execute inside containers as the container `Entrypoint`:

{{<highlight yaml "linenos=table,linenostart=5,hl_lines=2-5" >}}
steps:
- name: test
  image: golang:1.12
  commands:
  - go test

- name: build
  commands:
  - docker build .
{{< / highlight >}}

Or pipeline steps can execute directly on the host machine:

{{<highlight yaml "linenos=table,linenostart=5,hl_lines=7-9" >}}
steps:
- name: test
  image: golang:1.12
  commands:
  - go test

- name: build
  commands:
  - docker build .
{{< / highlight >}}
