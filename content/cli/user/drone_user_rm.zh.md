+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user rm"
url = "zh/cli-user-rm"
weight = 3

[menu.cli]
  weight = 42
  identifier = "cli-user-rm-zh"
  parent = "cli_user"
+++

<!--This subcommand deletes a registered user from the system. Please note this command requires administrative privileges.-->

这个子命令用来从系统中删除一个注册用户，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/user/data/drone_user_rm.out.txt" >}}
```

使用示例：

```text
$ drone user rm octocat
```
