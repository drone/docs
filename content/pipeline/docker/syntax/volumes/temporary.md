---
date: 2000-01-01T00:00:00+00:00
title: Temporary Volumes
author: bradrydzewski
weight: 19
toc: false
description: |
  Mount temporary volumes to share state between pipeline steps.
---

Temporary mounts are docker volumes that are created before the pipeline starts and destroyed when the pipeline completes. This can be used to share files or folders among pipeline steps.

```yaml {linenos=table, hl_lines=["8-10", "17-19", "23-25"]}
kind: pipeline
type: docker
name: default

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
```


# Tmpfs Volumes

Temporary mounts can also be configured to use tmpfs which mounts the volume in-memory. _Please note that you cannot share tmpfs volumes between containers; a tmpfs volume can only be mounted to a single container at a time._

```yaml {linenos=table, hl_lines=["18-19"]}
kind: pipeline
type: docker
name: default

steps:
- name: test
  image: golang
  volumes:
  - name: cache
    path: /go
  commands:
  - go get
  - go test
  - go build

volumes:
- name: cache
  temp:
    medium: memory
```
