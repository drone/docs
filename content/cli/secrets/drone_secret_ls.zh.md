+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret ls"
url = "zh/cli-secret-ls"

[menu.cli]
  weight = 14
  identifier = "cli-secret-ls-zh"
  parent = "cli_secret"
+++

<!--This subcommand returns a list of secrets for the named repository. Please note this command requires administrative privilege to the repository.-->

这个子命令返回对应仓库仓库密文，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/secrets/data/drone_secret_ls.out.txt" >}}
```

使用示例：

```text
$ drone repo ls -repository octocat/hello-world

docker_username
Events: push, tag, deployment
Images: plugins/docker

docker_password
Events: push, tag, deployment
Images: plugins/docker
```
