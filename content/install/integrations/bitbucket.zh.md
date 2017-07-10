+++
date = "2017-04-15T14:39:04+02:00"
title = "Bitbucket Cloud"
url = "zh/install-for-bitbucket-cloud"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-bitucket-cloud-zh"
  weight = 5
+++


<!--Drone comes with built-in support for Bitbucket Cloud. To enable Bitbucket Cloud you should configure the Drone container using the following environment variables:-->

Drone 支持 Bitbucket Cloud，使用下面的环境变量来配置使用 Bitbucket Cloud

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_BITBUCKET=true
+     - DRONE_BITBUCKET_CLIENT=95c0282573633eb25e82
+     - DRONE_BITBUCKET_SECRET=30f5064039e6b359e075
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

<!--# Configuration-->

# 配置

<!--This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.-->

下面是所有的配置选项。一般来说，使用默认配置可以满足绝大部分的安装需求：

DRONE_BITBUCKET=true 
: true 使用 Bitbucket

<!--: Set to true to enable the Bitbucket driver.-->-->

DRONE_BITBUCKET_CLIENT
: Bitbucket oauth2 client id

DRONE_BITBUCKET_SECRET
: Bitbucket oauth2 client secret

<!--# Registration-->

# 注册应用程序

<!--You must register your application with Bitbucket in order to generate a client and secret. Navigate to your account settings and choose OAuth from the menu, and click Add Consumer.-->

在 Bitbucket 上注册一个应用，并生成一个客户端和密钥。访问 Bitbucket 设置（account settings）页面，选择 OAuth 页面，点击 Add Consumer 。

<!--Please use the Authorization callback URL:-->

使用下面的认证回调 URL（Authorization callback URL），请修改域名为自定义域名：

```nohighlight
http://drone.mycompany.com/authorize
```

<!--Please also be sure to check the following permissions:-->

请确定确认允许使用下面的权限：

```nohighlight
Account:Email
Account:Read
Team Membership:Read
Repositories:Read
Webhooks:Read and Write
```

<!--# Missing Features-->

# 暂缺的特性

<!--Merge requests and mercurial repositories are not currently supported. We are interested in patches to include this functionality. If you are interested in contributing to Drone and submitting a patch please [contact us](https://gitter.im/drone/drone).-->

目前不支持 合并请求（Merge requests） and mercurial 仓库 。如果您对 Drone 感兴趣，并贡献这方面的补丁，[请在 Gitter 上与我们联系](https://gitter.im/drone/drone)。