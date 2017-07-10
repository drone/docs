+++
date = "2000-01-01T00:00:00+00:04"
title = "命令行 认证"
url = "zh/cli-authentication"

[menu.cli]
  weight = 2
  identifier = "cli-authentication-zh"
  parent = "cli_overview"
+++

<!--You will need to provide the Drone command line tools with your server address and personal access token. You can retrieve a Drone personal access token from your user profile screen. Click the show token button.-->

先设置 Drone 命令行工具的服务器地址和个人访问令牌（personal access token）。个人访问令牌可以从用户资料页面获得。点击 `show token` 按钮。

![user token](/images/drone_user_token.png)


<!--You can provide the server address and token using environment variables:-->

使用环境变量来提供服务器地址和令牌（token）。

```nohighlight
export DRONE_SERVER=http://drone.mycompany.com
export DRONE_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

<!--You can confirm the above configuration by running the below command:-->

使用下面的命令来确认上面的设置：

```nohighlight
$ drone info
User: octocat
Email: octocat@github.com
```
