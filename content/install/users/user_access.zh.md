+++
date = "2017-04-15T14:39:04+02:00"
title = "用户注册"
url = "zh/user-registration"

[menu.install]
  weight = 1
  identifier = "user-registration-zh"
  parent = "install_access"
+++

<!--Drone provides multiple configurations for open or limited access to the system. This section describes different options for user registration and access.-->

Drone 提供了多个配置项来开放或者限制系统的访问。这个部分描述了不同的用户注册和访问的选项。

<!--# Open Registration-->

# 开放注册

<!--Open registration is only recommended for secure installations on a private network. This configuration allows anyone to self-register and login to the system.-->

开放注册只推荐在私有网络上安全地安装。这个配置允许任何人自己注册和访问系统。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_OPEN= true
```

<!--# Restricted Registration-->

# 限制注册

<!--Restricted registration is the recommended configuration. This configuration allows members of white-listed organizations to self-register and login to the system.-->

限制注册是推荐的配置。这个配置允许白名单组织的成员注册和登录系统。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     DRONE_OPEN: true
+     DRONE_ORGS: dogpatch,dolores
```

<!--# Closed Registration-->

# 封闭注册

<!--Closed registration is enabled by default. Closed registration requires an administrator to manually register users using the command line utility. When using closed registration it is __imperative__ you provide a list of administrators that are able to login and manage accounts.-->

封闭注册是默认开启的模式。封闭注册需要管理员来手动使用命令行来添加用户。封闭注册需要提供一列管理员来管理用户注册。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
-     - DRONE_OPEN=true
+     - DRONE_OPEN=false
+     - DRONE_ADMIN=janedoe,johnsmith
```

<!--You can then manually grant users access using the command line utility:-->

设置了管理员后，管理员可以使用下面的命令来添加用户。

```nohighlight
drone user add <username>
```
