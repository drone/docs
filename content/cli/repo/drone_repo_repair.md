+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo repair"
url = "cli-repository-repair"

[menu.cli]
  weight = 3
  identifier = "cli-repository-repair"
  parent = "cli_repo"
+++

This subcommand re-creates webhooks for your repository in your version control system (e.g GitHub). This can be used if you accidentally delete your webhooks.

```text
{{< cat "content/cli/repo/data/drone_repo_repair.out.txt" >}}
```

Example usage:

```text
$ drone repo repair octocat/hello-world
```
