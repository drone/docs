+++
date = "2000-01-01T00:00:00+00:02"
title = "drone registry update"
url = "zh/cli-registry-update"

[menu.cli]
  weight = 21
  identifier = "cli-registry-update-zh"
  parent = "cli_registry"
+++

<!--This subcommand updates the named registry credentials. Please note this command requires authentication and administrative privilege to the repository.-->

这个子命令用来更新 registry 的登录凭据，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/registry/data/drone_registry_update.out.txt" >}}
```

Example usage:

```text
$ drone registry update \
  -repository octocat/hello-world \
  -hostname docker.io \
  -password ...
```
