---
date: 2022-17-01T00:00:00+00:00
title: ConfigMap Volumes
author: joekhoobyar
weight: 19
toc: false
description: |
  Mount configmap volumes to access kubernetes config maps in your pipeline.
---

Config Map volumes allow you to a kubernetes configmap as a volume into a pipeline step. This setting is only available to trusted repositories.

<div class="alert alert-warn">
This setting is only available to trusted repositories, since some kubernetes workloads (incorrectly) may choose to write sensitive date into a configmap - when it really should be written to a secret instead.
</div>

{{< highlight text "linenos=table,hl_lines=8-10 15-20" >}}
kind: pipeline
type: kubernetes
name: default

steps:
- name: build
  image: node
  volumes:
  - name: build-config
    path: /my-config-dir
  commands:
  - echo "an example of using configmaps - each key becomes a file"
  - echo "mykey=$(cat /my-config-dir/mykey)"

volumes:
- name: build-config
  config_map:
    name: my-build-config
    default_mode: 420     # same as 644 in octal, or u+w,a+r
    optional: false
{{< / highlight >}}

The first step is to define the configmap volume.  You can also assign a default mode, or even make the configmap optional.

{{< highlight text "linenos=table,linenostart=15" >}}
volumes:
- name: build-config
  config_map:
    name: my-build-config
    default_mode: 420     # same as 644 in octal, or u+w,a+r
    optional: false
{{< / highlight >}}

The next step is to configure your pipeline step to mount the volume into your container. The container path must also be an absolute path.  Each key in your configmap will become a file in the mounted directory.

{{< highlight text "linenos=table,linenostart=5" >}}
steps:
- name: build
  image: node
  volumes:
  - name: build-config
    path: /my-config-dir
{{< / highlight >}}
