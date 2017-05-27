+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo update"
url = "zh/cli-repository-update"

[menu.cli]
  weight = 4
  identifier = "cli-repository-update-zh"
  parent = "cli_repo"
+++

<!--This subcommand updates a named repository. Please note this command requires write access to the repository.-->

这个子命令用来更新一个仓库，这个子命令需要登录认证以及对应仓库的管理员权限。


```text
{{< cat "content/cli/repo/data/drone_repo_update.out.txt" >}}
```

<!--Example command updates the trusted flag:-->

更新 trusted 标签的示例：

```text
$ drone repo update octocat/hello-world --truted=true
```

<!--Example command updates the gated flag:-->

更新 gated 标签的示例：

```text
$ drone repo update octocat/hello-world --gated=true
```

<!--Example command updates the timeout value:-->

更新超时的示例：

```text
$ drone repo update octocat/hello-world --timeout=90m
```

<!--Example command updates the drone.yml file path-->

更新 drone.yml 的示例： 

```text
$ drone repo update octocat/hello-world --config=.github/.drone.yml
```
