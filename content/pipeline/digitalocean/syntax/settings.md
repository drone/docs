---
date: 2000-01-01T00:00:00+00:00
title: Settings
author: bradrydzewski
weight: -1
toc: false
description: |
  Configure the instance type.
---

Use the `settings` section to configure the Digital Ocean instance type. The runner creates the instance and executes pipeline commands on the instance using the ssh protocol.

{{< highlight text "linenos=table,hl_lines=8-11" >}}
kind: pipeline
type: digitalocean
name: default

token:
  from_secret: token

server:
  image: docker-18-04
  size: s-1vcpu-1gb
  region: nyc1

steps:
- name: build
  commands:
  - go build
  - go test
{{< / highlight >}}

# Digital Ocean Token

In the above example we provide the digital ocean token, sourced from a secret. This token is require in order to authenticate with the Digital Ocean API and create the server instance.

{{< highlight text "linenos=table,linenostart=5,hl_lines=1-3" >}}
token:
  from_secret: token

server:
  image: docker-18-04
  size: s-1vcpu-1gb
  region: nyc1

steps:
- name: build
  commands:
  - go build
  - go test
{{< / highlight >}}
