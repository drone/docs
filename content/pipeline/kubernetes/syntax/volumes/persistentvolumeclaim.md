---
date: 2022-17-01T00:00:00+00:00
title: PersistentVolumeClaim
author: joekhoobyar
weight: 19
toc: false
description: |
  Mount external volumes in your pipeline.
---

PersistentVolumeClaim allow you to mount exteral volumes into a pipeline step. 

```yaml {linenos=table, hl_lines=["8-10", "15-20"]}
kind: pipeline
type: kubernetes
name: default

steps:
- name: shell
  image: debian
  volumes:
  - name: shared
    path: /shared
  commands:
  - df -h

volumes:
- name: shared
  claim:
    name: received-data-claim
    read_only: false # <true|false>
```

The first step is to define the external volume. 

```yaml {linenos=table, linenostart=15}
volumes:
- name: cache
  claim:
    name: received-data-claim
    read_only: false # <true|false>
```

The next step is to configure your pipeline step to mount the named external volume path into your container. The container path must also be an absolute path.

```yaml {linenos=table, linenostart=5}
steps:
- name: shell
  image: debian
  volumes:
  - name: shared
    path: /shared
```
