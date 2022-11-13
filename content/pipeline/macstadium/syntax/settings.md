---
date: 2000-01-01T00:00:00+00:00
title: Settings
author: bradrydzewski
weight: -1
toc: false
draft: true
description: |
  Configure the instance type.
---

Use the `settings` section to configure the MacStadium virtual machine. The runner creates the virtual machine and executes pipeline commands on the instance using the ssh protocol.

```yaml {linenos=table, hl_lines=["5-7"]}
kind: pipeline
type: macstadium
name: default

settings:
  image: Drone.img
  cpu: 12

steps:
- name: build
  commands:
  - go build
  - go test
```
