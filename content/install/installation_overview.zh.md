+++
date = "2017-04-15T14:39:04+02:00"
title = "安装概述"
url = "zh/installation"

[menu.install]
  weight = 1
  identifier = "installation-zh"
  parent = "install_overview"
+++

<!--Drone is a lightweight, powerful continuous delivery platform built for containers. Drone is packaged and distributed as a Docker image and can be downloaded from Dockerhub.-->

Drone 是一个轻量级，为容器构建的强大的持续交付平台。 Drone 以 Docker 镜像的方式在 Dockerhub 上打包和发布。

```text
docker pull drone/drone:{{% version %}}
```

<!--# Docker Compose-->

# Docker Compose

<!--This section provides basic instructions for installing Drone using [docker-compose](https://docs.docker.com/compose/). The below configuration can be used to start the Drone server with a single agent.-->

这个章节提供了使用 [docker-compose](https://docs.docker.com/compose/) 安装 Drone 的基本步骤。下面的配置开启一个使用一个代理客户端（Agent）的 Drone 服务器：

```yaml
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}

  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

<!--Drone integrates with multiple version control providers, configured using environment variables. This example demonstrates basic GitHub integration.-->

Drone 可以整合多个版本控制系统，使用环境变量来配置。这个例子展示了 GitHub 的整合。

<!--You must register Drone with GitHub to obtain the client and secret. The authorization callback url must match `<scheme>://<host>/authorize`-->

{{% alert %}}
你需要在 GitHub 上注册 Drone 来获得 client 和 secret。认证回调URL （authorization callback url）应该填写 `<scheme>://<host>/authorize`
{{% /alert %}}

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_ORGS=dolores,dogpatch
      - DRONE_ADMIN=johnsmith,janedoe
+     - DRONE_GITHUB=true
+     - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
+     - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

<!--Drone mounts a volume on the host machine to persist the sqlite database.-->

Drone 在宿主机器上绑定了空间来持久化 sqlite 数据库。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
+   volumes:
+     - ./drone:/var/lib/drone/
    restart: always
```

<!--Drone needs to know its own address. You must therefore provide the address in `<scheme>://<hostname>` format. Please omit trailing slashes.-->

Drone 需要知道它自己的地址。你因此必须提供 `<scheme>://<hostname>` 的地址。请忽略末端的 `/`。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
+     - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

<!--Drone agents require access to the host machine Docker daemon.-->

Drone 代理客户端需要访问宿主 Docker daemon。

```diff
services:
  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on: [ drone-server ]
+   volumes:
+     - /var/run/docker.sock:/var/run/docker.sock
```

<!--Drone agents require the server address for agent-to-server communication.-->

Drone 代理服务器需要服务器地址来进行代理客户端和服务器的交流。

```diff
services:
  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on: [ drone-server ]
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
+     - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

<!--Drone server and agents use a shared secret to authenticate communication. This should be a random string of your choosing and should be kept private.-->

Drone 服务器和代理客户端使用一个共享密钥来认证交流。这个应为你选择一个随机的字符串，并不要公开。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
+     - DRONE_SECRET=${DRONE_SECRET}
  drone-agent:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_DEBUG: true
+     - DRONE_SECRET=${DRONE_SECRET}
```

<!--Drone registration is closed by default. This example enables open registration for users that are members of approved GitHub organizations.-->

Drone 默认是关闭注册的。这个例子开放了指定 GitHub 组织的成员的开放注册：

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_OPEN=true
+     - DRONE_ORGS=dolores,dogpatch
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

<!--Drone administrators should also be enumerated in your configuration.-->

Drone 管理员需要在配置文件中设置：

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_ORGS=dolores,dogpatch
+     - DRONE_ADMIN=johnsmith,janedoe
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```
