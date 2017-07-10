+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret rm"
url = "zh/cli-secret-rm"

[menu.cli]
  weight = 13
  identifier = "cli-secret-rm-zh"
  parent = "cli_secret"
+++

<!--This subcommand deletes a named repository secret. Please note this command requires administrative privilege to the repository.-->

这个子命令用来删除仓库的一个的密文，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/secrets/data/drone_secret_rm.out.txt" >}}
```

使用示例：


```text
$ drone secret rm -repository octocat/hello-world -name docker_password
```
