---
date: 2000-01-01T00:00:00+00:00
title: Volumes
author: tphoney
weight: 90
toc: true
draft: true
hide: true

description: |
  Mount host volumes to access the host filesystem in your pipeline.
---

# Host Volumes

Host mounts allow you to mount an absolute path on the host machine into a pipeline step. This setting is only available to trusted repositories.

<div class="alert alert-warn">
This setting is only available to trusted repositories, since mounting host machine volumes is a security risk.
</div>

{{<highlight yaml "linenos=table,hl_lines=8-10 15-18" >}}
kind: pipeline
type: aws
name: default

pool:
  use: ubuntu

steps:
- name: build
  image: node
  volumes:
  - name: cache
    path: /tmp/cache
  commands:
  - npm install
  - npm test

volumes:
- name: cache
  host:
    path: /var/lib/cache
{{< / highlight >}}

The first step is to define the host machine volume path. The host volume path must be an absolute path.

{{<highlight yaml "linenos=table,linenostart=15" >}}
volumes:
- name: cache
  host:
    path: /var/lib/cache
{{< / highlight >}}

The next step is to configure your pipeline step to mount the named host path into your container. The container path must also be an absolute path.

{{<highlight yaml "linenos=table,linenostart=5" >}}
steps:
- name: build
  image: node
  volumes:
  - name: cache
    path: /tmp/cache
{{< / highlight >}}

# Temporary Volumes

Temporary mounts are docker volumes that are created before the pipeline stars and destroyed when the pipeline completes. The can be used to share files or folders among pipeline steps.

{{<highlight yaml "linenos=table,hl_lines=8-10 17-19 23-25" >}}
kind: pipeline
type: aws
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
