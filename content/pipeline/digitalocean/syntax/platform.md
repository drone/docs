---
date: 2000-01-01T00:00:00+00:00
title: Platform
author: bradrydzewski
weight: 2
toc: false
description: |
  Configure the target operating system and architecture.
---

Use the `platform` section to configure the target operating system and architecture.

Example linux pipeline:

```yaml {linenos=table, hl_lines=["8-10"]}
kind: pipeline
type: digitalocean
name: default

token:
  from_secret: token

platform:
  os: linux
  arch: amd64

steps:
- name: build
  commands:
  - go build
  - go test
```

Example freebsd pipeline:

```yaml {linenos=table, hl_lines=["8-10"]}
kind: pipeline
type: digitalocean
name: default

token:
  from_secret: token

platform:
  os: freebsd
  arch: amd64

steps:
- name: build
  commands:
  - go build
  - go test
```