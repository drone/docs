+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo ls"
url = "zh/cli-repository-list"

[menu.cli]
  weight = 8
  identifier = "cli-repository-list-zh"
  parent = "cli_repo"
+++

<!--This subcommand returns a full list of user repositories.-->

这个子命令返回用户仓库列表。

```text
{{< cat "content/cli/repo/data/drone_repo_ls.out.txt" >}}
```

使用示例：

```text
$ drone repo ls

octocat/hello-world
octocat/fork-spoon
github/gitignore
github/hubot
github/hub
```

<!--Filter the output by organization:-->

按照组织来筛选：

```text
$ drone repo ls --org=github

github/gitignore
github/hubot
github/hub
```

<!--Format the output using a custom Go template:-->

使用 Go 模版来格式化输出：

```text
$ drone repo ls --format="{{ .Link }}"

https://github.com/octocat/hello-world
https://github.com/octocat/fork-spoon
https://github.com/github/gitignore
https://github.com/github/hubot
https://github.com/github/hub
```
