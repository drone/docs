+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo repair"
url = "zh/cli-repository-repair"

[menu.cli]
  weight = 3
  identifier = "cli-repository-repair-zh"
  parent = "cli_repo"
+++

<!--This subcommand re-creates webhooks for your repository in your version control system (e.g GitHub). This can be used if you accidentally delete your webhooks.-->

这个子命令重建版本控制系统（如 GitHub）的 webhooks，可以在意外地删除了 wenhooks 的情况下使用。

```text
{{< cat "content/cli/repo/data/drone_repo_repair.out.txt" >}}
```

使用示例：

```text
$ drone repo repair octocat/hello-world
```
