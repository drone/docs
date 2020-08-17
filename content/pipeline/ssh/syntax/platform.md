---
date: 2000-01-01T00:00:00+00:00
title: Platform
author: bradrydzewski
weight: 2
toc: false
description: |
  Configure the target operating system and architecture.
---

Use the `platform` section to configure the target operating system and architecture and route the pipeline to the appropriate runner.

Example windows pipeline:

{{< highlight text "linenos=table,hl_lines=10-13" >}}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

platform:
  os: windows
  arch: amd64

steps:
- name: build
  commands:
  - go build
  - go test
{{< / highlight >}}

# Supported Platforms

os          | arch
------------|-----
`linux`     | `amd64`
`linux`     | `arm64`
`linux`     | `arm`
`linux`     | `386`
`windows`   | `amd64`
`windows`   | `386`
`darwin`    | `amd64`
`freebsd`   | `amd64`
`freebsd`   | `arm`
`freebsd`   | `386`
`netbsd`    | `amd64`
`netbsd`    | `arm`
`openbsd`   | `amd64`
`openbsd`   | `arm`
`openbsd`   | `386`
`dragonfly` | `amd64`
`solaris`   | `amd64`
