+++
date = "2017-04-15T14:39:04+02:00"
title = "用户管理"
url = "zh/user-management"

[menu.install]
  weight = 2
  identifier = "user-management-zh"
  parent = "install_access"
+++

<!--You can manually manage user registration using the command line utility. Please see the command line documentation for installing and configuring the command line utility.-->

你可以使用命令行工具手动管理用户注册。请参考命令行文档来安装和配置命令行工具。

<!--Use the `ls` command to list all active users:-->

使用 `ls` 命令来列出所有活动用户：

```nohighlight
drone user ls
```

<!--Use the `add` command to add users to the system by login:-->

使用 `add` 命令来添加登录用户：

```nohighlight
drone user add octocat
```

<!--Use the `rm` command to remove users from the system by login:-->

使用 `rm` 命令来从系统删除用户：

```nohighlight
drone user rm octocat
```

<!--Please note that only drone administrators can manage users. Please see the below example to configure one or many administrators, separated by a comma, using the designated environment variable.-->

注意只有 Drone 管理员可以管理用户。参考下面的例子来使用环境变量配置一个或者多个管理员，以英文逗号分隔。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_ADMIN=janedoe,johnsmith
```
