---
date: 2000-01-01T00:00:00+00:00
title: Temporary Volumes
author: tphoney
weight: 19
toc: false
description: |
  Mount temporary volumes to share state between pipeline steps.
---

Temporary mounts are docker volumes that are created before the pipeline starts and destroyed when the pipeline completes. This can be used to share files or folders among pipeline steps.

{{<highlight yaml "linenos=table,hl_lines=8-10 17-19 23-25" >}}
kind: pipeline
type: vm
name: default

pool:
  use: ubuntu

steps:
- name: test
  image: golang
  volumes:
  - name: cache
    path: /go
  commands:
  - go get
  - go test

- name: build
  image: golang
  volumes:
  - name: cache
    path: /go
  commands:
  - go build

volumes:
- name: cache
  temp: {}
{{< / highlight >}}
