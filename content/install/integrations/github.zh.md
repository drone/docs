+++
date = "2017-04-15T14:39:04+02:00"
title = "GitHub"
url = "zh/install-for-github"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-github-zh"
  weight = 1
+++

<!--Drone comes with built-in support for GitHub and GitHub Enterprise. To enable GitHub you should configure the Drone container using the following environment variables:-->

Drone 支持 GitHub，，使用下面的环境变量来配置使用 GitHub。

```diff
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
+     - DRONE_GITHUB=true
+     - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
+     - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
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

<!--# Registration-->

# 注册应用程序

<!--Register your application with GitHub to create your client id and secret. It is very import the authorization callback URL matches your http(s) scheme and hostname exactly with `/authorize` as the path.-->

在 GitHub 注册应用程序来获得 client id 和 secret。认证回调 URL（authorization callback URL）应与您的对应域名匹配，并以 `/authorize` 作为 URL 的路径（path）。

<!--Please use this screenshot for reference:-->

请以这个截屏为参考：

![github oauth setup](images/github_oauth.png)

<!--# Configuration-->

# 配置

<!--This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.-->

下面是所有的配置选项。一般来说，使用默认配置可以满足绝大部分的安装需求：


<!--DRONE_GITHUB=true
: Set to true to enable the GitHub driver.

DRONE_GITHUB_URL=`https://github.com`
: GitHub server address.

DRONE_GITHUB_CLIENT
: Github oauth2 client id.

DRONE_GITHUB_SECRET
: Github oauth2 client secret.

DRONE_GITHUB_SCOPE=repo,repo:status,user:email,read:org
: Comma-separated Github oauth scope.

DRONE_GITHUB_GIT_USERNAME
: Optional. Use a single machine account username to clone all repositories.

DRONE_GITHUB_GIT_PASSWORD
: Optional. Use a single machine account password to clone all repositories.

DRONE_GITHUB_PRIVATE_MODE=false
: Set to true if Github is running in private mode.

DRONE_GITHUB_MERGE_REF=true
: Set to true to use the `refs/pulls/%d/merge` vs `refs/pulls/%d/head`

DRONE_GITHUB_CONTEXT=continuous-integration/drone
: Customize the GitHub status message context

DRONE_GITHUB_SKIP_VERIFY=false
: Set to true to disable SSL verification.-->

DRONE_GITHUB=true
: true 使用 GitHub

DRONE_GITHUB_URL=`https://github.com`
: GitHub server 地址

DRONE_GITHUB_CLIENT
: Github oauth2 client id

DRONE_GITHUB_SECRET
: Github oauth2 client secret

DRONE_GITHUB_SCOPE=repo,repo:status,user:email,read:org
: Github oauth scope，使用英文逗号 `,` 分隔

DRONE_GITHUB_GIT_USERNAME
: 可选，使用单一用户来克隆所有仓库，这个用户的用户名

DRONE_GITHUB_GIT_PASSWORD
: 可选，使用单一用户来克隆所有仓库，这个用户的密码

DRONE_GITHUB_PRIVATE_MODE=false
: 如果 GitHub 以 private 私有模式运行，应设置为 true 

DRONE_GITHUB_MERGE_REF=true
: 设置为 true 来使用 `refs/pulls/%d/merge` ，反之则使用 `refs/pulls/%d/head`

DRONE_GITHUB_CONTEXT=continuous-integration/drone
: 自定义 GitHub status message 上下文

DRONE_GITHUB_SKIP_VERIFY=false
: 设置 true 来取消 SSL 检查
