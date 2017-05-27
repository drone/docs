+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo rm"
url = "zh/cli-repository-remove"

[menu.cli]
  weight = 9
  identifier = "cli-repository-remove-zh"
  parent = "cli_repo"
+++

<!--This subcommand deletes the named repository from the system. Please note this command requires administrative access to the repository.-->

这个子命令用来从 Drone 删除一个仓库，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/repo/data/drone_repo_rm.out.txt" >}}
```

使用示例：

```text
$ drone repo rm octocat/hello-world
```
