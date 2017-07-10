+++
date = "2000-01-01T00:00:00+00:02"
title = "drone registry rm"
url = "zh/cli-registry-rm"

[menu.cli]
  weight = 23
  identifier = "cli-registry-rm-zh"
  parent = "cli_registry"
+++

<!--This subcommand deletes a named registry credentials. Please note this command requires authentication and administrative privilege to the repository.-->

这个子命令用来删除一个 registry 的登录凭据，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/registry/data/drone_registry_rm.out.txt" >}}
```

<!--Example usage:-->

使用示例：

```text
$ drone registry rm -repository octocat/hello-world
```

<!--Example usage with custom docker registry:-->

自定义 Docker registry 示例：

```text
$ drone registry rm \
  -repository octocat/hello-world \
  -hostname gcr.io
```
