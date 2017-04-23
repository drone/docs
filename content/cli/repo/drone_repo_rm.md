+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo rm"
url = "cli-repository-remove"

[menu.cli]
  weight = 9
  identifier = "cli-repository-remove"
  parent = "repo"
+++

This subcommand deletes the named repository from the system. Please note this command requires administrative access to the repository.

```text
{{< cat "content/cli/repo/data/drone_repo_rm.out.txt" >}}
```

Example usage:

```text
$ drone repo rm octocat/hello-world
```
