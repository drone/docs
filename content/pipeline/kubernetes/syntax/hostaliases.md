---
date: 2000-01-01T00:00:00+00:00
title: HostAliases
author: ysicing
weight: 8
toc: false
description: |
  add additional entries to the hosts file
---

The `HostAliases` section can add entries to a Pod's `/etc/hosts` file provides Pod-level override of hostname resolution when DNS and other options are not applicable

{{< link-to "https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/" >}}
Learn more about HostAliases
{{</ link-to >}}

Example configuration:

{{< highlight text "linenos=table,hl_lines=11-15" >}}
kind: pipeline
type: kubernetes
name: default

steps:
- name: cat-hosts
  image: busybox
  commands:
  - cat /etc/hosts

host_aliases:
  - ip: 127.0.0.1
    hostnames:
      - www.baidu.com
{{< / highlight >}}
