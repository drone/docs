+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user add"
url = "zh/cli-user-add"
weight = 1

[menu.cli]
  weight = 40
  identifier = "cli-user-add-zh"
  parent = "cli_user"
+++

<!--This subcommand registers a new user with the system. Please note this command requires administrative privileges.-->

这个子命令用来注册一个新用户，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/user/data/drone_user_add.out.txt" >}}
```

<!--Example usage, adds a user by username:-->

添加一个指定用户名的用户：

```text
$ drone user add octocat
```
