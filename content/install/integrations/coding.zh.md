+++
date = "2017-07-23T16:51:55+08:00"
title = "Coding"
url = "zh/install-for-coding"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-coding-zh"
  weight = 7
+++

Drone 内置支持 Coding。请使用下列的环境变量来配置 Drone 容器以启用 Coding：

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_CODING=true
+     - DRONE_CODING_CLIENT=${DRONE_CODING_CLIENT}
+     - DRONE_CODING_SECRET=${DRONE_CODING_SECRET}
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

# 注册应用

注册 Coding 应用以获得 Client ID 和 Client Secret。进入『账户』的『应用管理』页面，选择『添加应用』便可看到『应用设置』页面。『回调地址』填写与『应用主页』相对应的主机名，并以 `/authorize` 作为 URL 的路径。

<a href="images/coding_oauth.png" target="_blank"><img src="images/coding_oauth.png" alt="coding oauth setup"></a>

# 配置

下列是完整的配置选项。其中的大部分选项使用默认值便可满足大多数的安装需求。

DRONE_CODING=true
: 设置为 true 以启用 Coding。

DRONE_CODING_URL=`https://coding.net`
: Coding 服务器 URL。默认值使用的是平台版的 URL。企业版应设置为 `https://{company}.coding.net`。

DRONE_CODING_CLIENT
: Coding 应用的 Client ID。

DRONE_CODING_SECRET
: Coding 应用的 Client Secret.

DRONE_CODING_SCOPE=user,project,project:depot
: Coding 应用用户数据访问权限，用英文逗号『,』隔开。

DRONE_CODING_GIT_MACHINE=git.coding.net
: Coding git 服务器主机名。默认值使用的是平台版的主机名。企业版应设置为 `e.coding.net`。

DRONE_CODING_GIT_USERNAME
: 可选。用于克隆所有仓库的用户名。

DRONE_CODING_GIT_PASSWORD
: 可选。用于克隆所有仓库的密码.

DRONE_CODING_SKIP_VERIFY=false
: 设置为 true 以禁止 SSL 证书校验。
