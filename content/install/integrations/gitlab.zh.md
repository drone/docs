+++
date = "2017-04-15T14:39:04+02:00"
title = "GitLab"
url = "zh/install-for-gitlab"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-gitlab-zh"
  weight = 2
+++

<!--Drone comes with built-in support for the GitLab version 8.2 and higher. To enable GitLab you should configure the Drone container using the following environment variables:-->

Drone 支持 GitLab，使用下面的环境变量来配置使用 GitLab。

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_GITLAB=true
+     - DRONE_GITLAB_CLIENT=95c0282573633eb25e82
+     - DRONE_GITLAB_SECRET=30f5064039e6b359e075
+     - DRONE_GITLAB_URL=http://gitlab.mycompany.com
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


<!--DRONE_GITLAB=true
: Set to true to enable the GitLab driver.

DRONE_GITLAB_URL=`https://gitlab.com`
: GitLab Server address.

DRONE_GITLAB_CLIENT
: GitLab oauth2 client id.

DRONE_GITLAB_SECRET
: GitLab oauth2 client secret.

DRONE_GITLAB_GIT_USERNAME
: Optional. Use a single machine account username to clone all repositories.

DRONE_GITLAB_GIT_PASSWORD
: Optional. Use a single machine account password to clone all repositories.

DRONE_GITLAB_SKIP_VERIFY=false
: Set to true to disable SSL verification.

DRONE_GITLAB_PRIVATE_MODE=false
: Set to true if GitLab is running in private mode.-->


DRONE_GITLAB=true
: true 使用 GitLab

DRONE_GITLAB_URL=`https://gitlab.com`
: GitLab Server 地址

DRONE_GITLAB_CLIENT
: GitLab oauth2 client id

DRONE_GITLAB_SECRET
: GitLab oauth2 client secret

DRONE_GITLAB_GIT_USERNAME
: 可选，使用单一用户来克隆所有仓库，这个用户的用户名

DRONE_GITLAB_GIT_PASSWORD
: 可选，使用单一用户来克隆所有仓库，这个用户的密码

DRONE_GITLAB_SKIP_VERIFY=false
: 设置 true 来取消 SSL 检查

DRONE_GITLAB_PRIVATE_MODE=false
: 如果 GitLab 以 private 私有模式运行，应设置为 true 

<!--# Registration-->

# 注册应用程序


<!--You must register your application with GitLab in order to generate a Client and Secret. Navigate to your account settings and choose Applications from the menu, and click New Application.-->

在 GitLab 上注册一个应用，并生成一个客户端和密钥。访问账户设置（account settings）页面，选择 Applications 页面，点击 New Application 。

使用下面的认证回调 URL（Authorization callback URL），请修改域名为自定义域名： `http://drone.mycompany.com/authorize`

<!--Please use `http://drone.mycompany.com/authorize` as the Authorization callback URL.-->
