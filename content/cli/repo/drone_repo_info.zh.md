+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo info"
url = "zh/cli-repository-info"

[menu.cli]
  weight = 5
  identifier = "cli-repository-info-zh"
  parent = "cli_repo"
+++

<!--This subcommand prints the named repository details.-->

这个子命令打印一个仓库的详细信息。

```text
{{< cat "content/cli/repo/data/drone_repo_info.out.txt" >}}
```

使用示例：

```text
$ drone repo info octocat/hello-world

Owner: octocat
Repo: hello-world
Type: git
Private: false
Remote: https://github.com/octocat/hello-world.git
```

<!--Format the output using a custom Go template:-->

使用 Go 模版来格式化输出：

```text
$ drone repo info --format="{{ .Link }}" octocat/hello-world

https://github.com/octocat/hello-world.git
```
