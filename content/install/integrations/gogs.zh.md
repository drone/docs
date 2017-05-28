+++
date = "2017-04-15T14:39:04+02:00"
title = "Gogs"
url = "zh/install-for-gogs"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-gogs-zh"
  weight = 3
+++

<!--Drone comes with built-in support for the latest stable version of Gogs. To enable Gogs you should configure the Drone container using the following environment variables:-->

Drone 支持 Gogs，使用下面的环境变量来配置使用 Gogs。

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_GOGS=true
+     - DRONE_GOGS_URL=http://gogs.mycompany.com
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

<!--# Authentication-->

# 认证

<!--Drone will prompt you for a username and password to authenticate. You should use your Gogs username and password. This is unfortunately required due to Gogs lack of oauth2 support.-->

Gogs 不支持 oauth2 认证，Drone 会提示输入用户名和密码来进行认证。请使用 Gogs 的用户名和密码。

<!--# Configuration-->

# 配置

<!--This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.-->

下面是所有的配置选项。一般来说，使用默认配置可以满足绝大部分的安装需求：


<!--DRONE_GOGS=true
: Set to true to enable the Gogs driver.

DRONE_GOGS_URL
: Gogs server address.

DRONE_GOGS_GIT_USERNAME
: Optional. Use a single machine account username to clone all repositories.

DRONE_GOGS_GIT_PASSWORD
: Optional. Use a single machine account password to clone all repositories.

DRONE_GOGS_PRIVATE_MODE=false
: Set to true if Gogs is running in private mode.

DRONE_GOGS_SKIP_VERIFY=false
: Set to true to disable SSL verification.-->

DRONE_GOGS=true
: true 使用 Gogs

DRONE_GOGS_URL
: Gogs server 地址

DRONE_GOGS_GIT_USERNAME
: 可选，使用单一用户来克隆所有仓库。这个用户的用户名

DRONE_GOGS_GIT_PASSWORD
: 可选，使用单一用户来克隆所有仓库。这个用户的密码

DRONE_GOGS_PRIVATE_MODE=false
: 如果 Gogs 以 private 私有模式运行，应设置为 true

DRONE_GOGS_SKIP_VERIFY=false
: 设置 true 来取消 SSL 检查
