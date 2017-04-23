+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo add"
url = "cli-repository-add"

[menu.cli]
  weight = 2
  identifier = "cli-repository-add"
  parent = "cli_repo"
+++

This subcommand registers a named repository with Drone. Please note this command requires administrative access to the repository.

```text
{{< cat "content/cli/repo/data/drone_repo_add.out.txt" >}}
```

Example usage:

```text
$ drone repo add octocat/hello-world
```
