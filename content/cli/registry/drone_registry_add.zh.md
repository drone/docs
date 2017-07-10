+++
date = "2000-01-01T00:00:00+00:02"
title = "drone registry add"
url = "zh/cli-registry-add"

[menu.cli]
  weight = 20
  identifier = "cli-registry-add-zh"
  parent = "cli_registry"
+++

<!--This subcommand adds registry credentials to your repository. Please note this command requires authentication and administrative privilege to the repository.-->

这个子命令用来添加 registry 的登录凭据，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/registry/data/drone_registry_add.out.txt" >}}
```

<!--Example usage:-->

使用示例：

```text
$ drone registry add \
  -repository octocat/hello-world \
  -username octocat \
  -password password
```

<!--Example loads the repository password from file:-->

从文件读取仓库密码的示例：

```text
$ drone registry add \
  -repository octocat/hello-world \
  -hostname gcr.io \
  -username octocat \
  -password @/path/to/token.js
```
