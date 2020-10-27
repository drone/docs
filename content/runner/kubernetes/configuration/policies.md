---
date: 2000-01-01T00:00:00+00:00
title: Policies
author: bradrydzewski
weight: 2
draft: true
description: |
  Configure policies to apply pipeline defaults.
---



```
---
kind: policy
name: octocat

match:
  repos:
  - "octocat/*"
  - "octocat/hello-world"

metadata:
  namespace: octocat

resources:
  request:
    cpu: 1
    memory: 512MiB
  limit:
    cpu: 4
    memory: 1GiB

node_selector:
  disktype: ssd

---
kind: policy
name: default

metadata:
  namespace: default
```


# Examples

```
---
kind: policy
name: default

service_account: drone
```


```
---
kind: policy
name: default

service_account: drone
```