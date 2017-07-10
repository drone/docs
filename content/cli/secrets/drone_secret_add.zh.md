+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret add"
url = "zh/cli-secret-add"

[menu.cli]
  weight = 10
  identifier = "cli-secret-add-zh"
  parent = "cli_secret"
+++

<!--This subcommand adds a secret to your repository secret store. Please note this command requires administrative privilege to the repository.-->

这个子命令给仓库添加密文，这个子命令需要登录认证以及对应仓库的管理员权限。

```text
{{< cat "content/cli/secrets/data/drone_secret_add.out.txt" >}}
```

使用示例：

```text
$ drone secret add \
  -repository octocat/hello-world \
  -name docker_password \
  -value ...
```

<!--Example usage limits the secret to a specific image:-->

限制只有指定镜像可以使用密文地示例：

```diff
$ drone secret add \
  -repository octocat/hello-world \
  -name docker_password \
+ -image plugins/docker \
  -value ...
```

<!--Example usage limits the secret to specific hook events:-->

限制只有指定 webhook 可以使用密文地示例：

```diff
$ drone secret add \
  -repository octocat/hello-world \
  -name docker_password \
  -image plugins/docker \
+ -event push \
+ -event pull_request \
+ -event tag \
+ -event deployment \
  -value ...
```

<!--Example usage adds the secret from a file:-->

从文件读取密文地示例：

```diff
$ drone secret add \
  -repository octocat/hello-world \
  -name ssh_key \
+ -value @/root/.ssh/id_rsa
```
