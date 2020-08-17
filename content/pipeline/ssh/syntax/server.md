---
date: 2000-01-01T00:00:00+00:00
title: Server
author: bradrydzewski
weight: -1
toc: true
description: |
  Configure the remote ssh server.
---

Use the `server` section to configure the remote ssh server. The runner connects to this server and executes pipeline commands using the ssh protocol.

{{< highlight text "linenos=table,hl_lines=5-9" >}}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password: correct-horse-battery-staple

steps:
- name: build
  commands:
  - go build
  - go test
{{< / highlight >}}

# Secrets

In the above example we hard-coded the server address and login credentials. For security reasons you should source these values from secrets.

{{< highlight text "linenos=table,linenostart=5,hl_lines=1-7" >}}
server:
  host:
    from_secret: host
  user:
    from_secret: username
  password:
    from_secret: password

steps:
- name: build
  commands:
  - go build
  - go test
{{< / highlight >}}

# SSH Keys

{{< highlight text "linenos=table,linenostart=5,hl_lines=7-8" >}}
server:
  host:
    from_secret: host
  user:
    from_secret: username
  ssh_key:
    from_secret: ssh_key
{{< / highlight >}}
