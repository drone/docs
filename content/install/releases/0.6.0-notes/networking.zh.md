+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 网络"
url = "zh/release-0.6.0-networking"
+++

<!--Drone now aligns with docker-compose and bridge networking with hostnames. This means services are no longer accessible on `localhost` or `127.0.0.1` and are instead available using their hostname as defined in the yaml configuration file.-->

Drone 现在支持 docker-compose 和 桥接网络的  hostnames，这意味着服务不在从 `localhost` 或者 `127.0.0.1` 来访问，而是从它们的 在 yaml 文件中定义的 hostname 来访问。

```diff
pipeline:
  build:
    image: golang
    commands:
-     - mysql -h localhost:3306
+     - mysql -h mysql:3306

services:
  mysql:
    image: mysql
```

<!--There are no immediate plans to continue support for pod networking (e.g. localhost) unless natively supported by Docker. I have opened [this issue](https://github.com/docker/docker/issues/26024) with the Docker project to request native pod networking. Please considering voicing your support if this capability is important to you or your organization.-->

目前没有继续支持 pod network 的计划，即 localhost，除非 Docker 支持这一网络。[这个 issue](https://github.com/docker/docker/issues/26024) 是有关 请求 Docker 项目原生支持 pod networking。如果您或您的组织有相应的需求，请反馈到这里。
