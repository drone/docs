---
date: 2000-01-01T00:00:00+00:00
title: Metadata
author: bradrydzewski
weight: -1
description: |
  Configure Kubernetes pipeline metadata.
---

Use the `metadata` section to provide uniquely identify pipeline resources. Example configuration defines the pipeline namespace:

```yaml {linenos=table, hl_lines=["5-6"]}
kind: pipeline
type: kubernetes
name: default

metadata:
  namespace: default

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test
```


Example with annotations:

```yaml {linenos=table, hl_lines=["7-9"]}
kind: pipeline
type: kubernetes
name: default

metadata:
  namespace: default
  annotations:
    key1: value1
    key2: value2
```

Example with labels:

```yaml {linenos=table, hl_lines=["7-9"]}
kind: pipeline
type: kubernetes
name: default

metadata:
  namespace: default
  labels:
    key1: value1
    key2: value2
```