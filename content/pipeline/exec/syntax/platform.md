---
date: 2000-01-01T00:00:00+00:00
title: Platform
author: bradrydzewski
weight: 1
toc: false
description: |
  Configure the target operating system and architecture.
---

Use the `platform` section to configure the target operating system and architecture and route the pipeline to the appropriate runner.

Example macOS (darwin) pipeline:

```yaml {linenos=table, hl_lines=["5-7"]}
kind: pipeline
type: exec
name: default

platform:
  os: darwin
  arch: amd64

steps:
- name: build
  commands:
  - go build
  - go test
```

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

