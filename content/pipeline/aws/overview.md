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

aliases:
- /yaml-reference/
- /user-guide/pipeline

description: |
  Overview of AWS pipelines.
---

## Overview

AWS pipelines can execute pipeline commands in two ways:

+ inside ephemeral Docker containers. Docker containers provide isolation, allowing safe execution of concurrent pipelines on the same machine.
+ run directly on the EC2 instance itself.

Example pipeline configuration:

{{< highlight text "linenos=table" >}}
---
kind: pipeline
type: aws
name: default

pool:
  use: ubuntu

steps:
- name: greeting
  image: golang:1.12
  commands:
  - go build
  - go test
- name: echo
  commands:
  - ls
...
{{< / highlight >}}

## Kind and Type

The first step is run in an ephemeral Docker container executing some go commands. The second step is run on the EC2 instance running the `ls` command.

The kind and type attributes define a aws pipeline.

{{< highlight text "linenos=table" >}}
---
kind: pipeline
type: aws
{{< / highlight >}}

## Pool

This section describes which AWS pool resources are used by the pipeline.

{{< highlight text "linenos=table" >}}
pool:
  use: ubuntu
{{< / highlight >}}

For more information about configuring pools, see  See [Pool file]({{< relref "../../runner/aws/configuration/pool.md" >}}) for more information.

## Steps

The `steps` section defines a series of shell commands. These commands are executed inside the Docker container as the `Entrypoint`. If any command returns a non-zero exit code, the pipeline fails and exits.

{{< highlight text "linenos=table,linenostart=6" >}}
steps:
- name: greeting
  image: golang:1.12
  commands:
  - go build
  - go test
{{< / highlight >}}
