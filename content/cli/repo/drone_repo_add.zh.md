+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo add"
url = "zh/cli-repository-add"

[menu.cli]
  weight = 2
  identifier = "cli-repository-add-zh"
  parent = "cli_repo"
+++

<!--This subcommand registers a named repository with Drone. Please note this command requires administrative access to the repository.-->

这个子命令用来添加一个仓库到 Drone，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/repo/data/drone_repo_add.out.txt" >}}
```

使用示例：

```text
$ drone repo add octocat/hello-world
```
