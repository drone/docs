+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret info"
url = "zh/cli-secret-info"

[menu.cli]
  weight = 12
  identifier = "cli-secret-info-zh"
  parent = "cli_secret"
+++

<!--This subcommand prints the named secret details. Please note this command requires administrative privilege to the repository.-->

这个子命令用来打印密文信息，这个子命令需要登录认证以及对应仓库的管理员权限。


```text
{{< cat "content/cli/secrets/data/drone_secret_info.out.txt" >}}
```

使用示例：

```text
$ drone secret info \
  -repository octocat/hello-world \
  -name docker_password

docker_password
Events: push, tag, deployment
Images: plugins/docker
```
