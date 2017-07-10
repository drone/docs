+++
date = "2017-04-15T14:39:04+02:00"
title = "Privileged Mode 特权模式"
url = "zh/privileged-mode"

[menu.usage]
  weight = 5
  identifier = "privileged-mode-zh"
  parent = "usage_config"
+++

<!--Drone gives the ability to configure privileged mode in the Yaml. You can use this parameter to launch containers with escalated capabilities.-->

可以通过配置 Yaml 文件来使用特权模式。可以使用 `privileged` 参数来启动具有提权能力的容器。

<!--Privileged mode is only available to trusted repositories and for security reasons should only be used in private environments.-->

{{% alert error %}}
特权模式只能在可信的仓库中使用。为了保证系统安全，请只在私有环境中使用特权模式。
{{% /alert %}}

```diff
pipeline:
  build:
    image: docker
    environment:
      - DOCKER_HOST=tcp://0.0.0.0:2375
    commands:
      - docker --tls=false ps

services:
  docker:
    image: docker:dind
    command: [ "--storage-driver=vfs", "--tls=false" ]
+   privileged: true
```
