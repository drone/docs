+++
date = "2017-04-15T14:39:04+02:00"
title = "Privileged Mode"
url = "privileged-mode"

[menu.usage]
  weight = 5
  identifier = "privileged-mode"
  parent = "usage_config"
+++

Drone gives the ability to configure privileged mode in the Yaml. You can use this parameter to launch containers with escalated capabilities.

{{% alert error %}}
Privileged mode is only available to trusted repositories and for security reasons should only be used in private environments.
{{% /alert %}}

```diff
pipeline:
  build:
    image: docker
    environment:
      - DOCKER_HOST=tcp://docker:2375
    commands:
      - docker --tls=false ps

services:
  docker:
    image: docker:dind
    command: [ "--storage-driver=vfs", "--tls=false" ]
+   privileged: true
```
