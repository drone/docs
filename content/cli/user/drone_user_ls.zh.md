+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user ls"
url = "zh/cli-user-ls"
weight = 4

[menu.cli]
  weight = 43
  identifier = "cli-user-ls-zh"
  parent = "cli_user"
+++

<!--This subcommand prints a list of all registered users. Please note this command requires administrative privileges.-->

这个子命令用来打印注册用户的列表，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/user/data/drone_user_ls.out.txt" >}}
```

使用示例：

```text
$ drone user ls

jonhsmith
janedoe
octocat
```

Format the output using a custom Go template:

```text
$ drone user ls --format="{{ .Login }} <{{ .Email }}>"

jonhsmith <jonh.smith@github.com>
janedoe   <jane.doe@github.com>
octocat   <octocat@github.com>
```
