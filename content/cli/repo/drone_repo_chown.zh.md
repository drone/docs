+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo chown"
url = "zh/cli-repository-chown"

[menu.cli]
  weight = 2
  identifier = "cli-repository-chown-zh"
  parent = "cli_repo"
+++

<!--This subcommand lets a user assume ownership of a named repository. Please note this command requires administrative access to the repository.-->

这个子命令用来让一个用户获得一个仓库到所有权，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/repo/data/drone_repo_chown.out.txt" >}}
```

<!--Below command assumes ownership to the named repository. Note that the individual executing this command becomes the owner.-->

下面的命令让运行命令的用户获得对应仓库到所有权。

```text
$ drone repo chown octocat/hello-world
```
