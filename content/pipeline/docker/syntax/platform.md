---
date: 2000-01-01T00:00:00+00:00
title: Platform
author: bradrydzewski
weight: 20
disable_toc: true
description: |
  Configure the target operating system and architecture.
---

Use the `platform` section to configure the target operating system and architecture and routes the pipeline to the appropriate runner. If unspecified, the system defaults to Linux amd64.

Example Linux amd64 pipeline:

{{< highlight yaml "linenos=table,hl_lines=5-7" >}}
kind: pipeline
type: docker
name: default

platform:
  os: linux
  arch: amd64

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}

Example Linux arm64 pipeline:

{{< highlight text "linenos=table,linenostart=5,hl_lines=3" >}}
platform:
  os: linux
  arch: arm64
{{< / highlight >}}

Example Linux arm32 pipeline:

{{< highlight text "linenos=table,linenostart=5,hl_lines=3" >}}
platform:
  os: linux
  arch: arm
{{< / highlight >}}

# Windows

If you are running Docker pipelines on windows you must specify the operating system version number.

<div class="alert">
Please note Windows pipelines are not available on Drone Cloud. This feature is only available when self-hosting
</div>

Example windows 1809 pipeline:

{{< highlight text "linenos=table,linenostart=5,hl_lines=4" >}}
platform:
  os: windows
  arch: amd64
  version: 1809
{{< / highlight >}}

Example windows 1903 pipeline:

{{< highlight text "linenos=table,linenostart=5,hl_lines=4" >}}
platform:
  os: windows
  arch: amd64
  version: 1903
{{< / highlight >}}

