+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user info"
url = "zh/cli-user-info"
weight = 2

[menu.cli]
  weight = 41
  identifier = "cli-user-info-zh"
  parent = "cli_user"
+++

<!--This subcommand prints information about the named registered user. Please note this command requires administrative privileges.-->

这个子命令用来打印一个注册用户的信息，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/user/data/drone_user_info.out.txt" >}}
```

使用示例：

```text
$ drone user info octocat

User: octocat
Email: octocat@github.com
```

<!--Format the output using a custom Go template:-->

使用 Go 模版来格式化输出：

```text
$ drone user info --format="{{ .Login }} <{{ .Email }}> octocat"

octocat <octocat@github.com>
```
