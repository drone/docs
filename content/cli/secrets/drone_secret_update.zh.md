+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret update"
url = "zh/cli-secret-update"

[menu.cli]
  weight = 11
  identifier = "cli-secret-update-zh"
  parent = "cli_secret"
+++

<!--This subcommand updates a secret in your repository secret store. Please note this command requires administrative privilege to the repository.-->

这个子命令用来更新仓库的一个密文，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/secrets/data/drone_secret_update.out.txt" >}}
```

使用示例：

```text
$ drone secret update \
  -repository octocat/hello-world \
  -name docker_username \
  -value ...
```
