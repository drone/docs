+++
date = "2017-04-15T14:39:04+02:00"
title = "用户管理员"
url = "zh/user-admins"

[menu.install]
  weight = 3
  identifier = "user-admins-zh"
  parent = "install_access"
+++

<!--You can grant administrative privileges to users by providing an enumerated user list, separated by a comma, using the designated environment variable.-->

使用环境变量以逗号分隔符来指定拥有管理员权限的用户。

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_ADMIN=janedoe,johnsmith
```

<!--Please note that the usernames are case-sensitive and must match the exact values as returned from your version control system (e.g. GitHub).-->

用户名是大小写敏感的，必须和您使用的版本控制系统（如 GitHub）返回的数值相匹配。
